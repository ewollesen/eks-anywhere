//go:build e2e
// +build e2e

package e2e

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/test/framework"
)

var manager testClusterProvider

type testClusterProvider interface {
	Setup() error
	Teardown()
	WithCluster(*testing.T, runFunc, cleanupFunc)
}

type vsphereClusterManager struct {
	mu      sync.Mutex
	e2eTest *framework.ClusterE2ETest
}

func (m *vsphereClusterManager) Setup() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	fmt.Printf("\n\n--- CLUSTER MANAGER SETUP ---\n\n")

	m.e2eTest = nil
	return nil
}

func (m *vsphereClusterManager) Teardown() {
	m.mu.Lock()
	defer m.mu.Unlock()

	// By definition the tests are completed at this point, so any logging will cause a panic

	m.e2eTest.T = &testing.T{}
	fmt.Printf("\n\n--- CLUSTER MANAGER TEARDOWN STARTING ---\n\n")
	m.e2eTest.StopIfFailed()
	m.e2eTest.DeleteCluster()
	m.e2eTest = nil
	fmt.Printf("\n\n--- CLUSTER MANAGER TEARDOWN COMPLETE ---\n\n")
}

type runFunc func(*framework.ClusterE2ETest)

type cleanupFunc func(*framework.ClusterE2ETest)

func alert(t *testing.T, msg string) {
	t.Helper()
	t.Logf("\n\n!!!\n!!! %s\n!!!\n\n", msg)
}

type reuseLogger struct {
	log *log.Logger
}

func newStderrLogger() *reuseLogger {
	return &reuseLogger{log.New(os.Stderr, "", 0)}
}

var _ framework.Logger = (*reuseLogger)(nil)

func (l *reuseLogger) Log(args ...any) {
	l.log.Print(args...)
}

func (l *reuseLogger) Logf(format string, args ...any) {
	l.log.Printf(format, args...)
}

func (m *vsphereClusterManager) WithCluster(t *testing.T, run runFunc, cleanup cleanupFunc) {
	m.mu.Lock()
	defer m.mu.Unlock()

	logger := newStderrLogger()
	if m.e2eTest == nil {
		alert(t, "Bringing up a cluster")
		test := framework.NewReusableClusterE2ETest(t, logger,
			framework.NewVSphere(t, framework.WithBottleRocket123()),
			framework.WithClusterFiller(api.WithKubernetesVersion(v1alpha1.Kube123)),
			framework.WithPackageConfig(t, packageBundleURI(v1alpha1.Kube121),
				EksaPackageControllerHelmChartName, EksaPackageControllerHelmURI,
				EksaPackageControllerHelmVersion, EksaPackageControllerHelmValues))
		test.GenerateClusterConfig()
		test.CreateCluster()
		m.e2eTest = test
	} else {
		alert(t, fmt.Sprintf("Re-using existing cluster and test cluster name: %q", m.e2eTest.ClusterName))
		m.e2eTest.T = t
	}

	defer func() {
		alert(t, "Cleaning up after a test")
		cleanup(m.e2eTest)
	}()
	run(m.e2eTest)
}

func init() {
	manager = &vsphereClusterManager{}
}

var _ testClusterProvider = (*vsphereClusterManager)(nil)

func TestMain(m *testing.M) {
	if err := manager.Setup(); err != nil {
		log.Fatal(err)
	}
	defer manager.Teardown()
	code := m.Run()
	fmt.Printf("\n\n ### Main test run complete ###\n\n")
	os.Exit(code)
}

func TestClusterReuse(s *testing.T) {
	s.Run("test number one", func(t *testing.T) {
		alert(t, "Starting first test")
		manager.WithCluster(s, runCuratedPackageInstallWithName("test1"), cleanupCuratedPackageInstall("test1"))
	})

	s.Run("test number two", func(t *testing.T) {
		alert(t, "Starting second test")
		manager.WithCluster(s, runCuratedPackageInstallWithName("test2"), cleanupCuratedPackageInstall("test2"))
	})
}

// Just for easy reference
// func runCuratedPackageInstallWithName(name string) func(*framework.ClusterE2ETest) {
// 	return func(test *framework.ClusterE2ETest) {
// 		test.T.Logf("\n\n!!!\n!!! runCuratedPackageInstall %q\n!!!\n\n", name)
// 		packageName := "hello-eks-anywhere"
// 		test.InstallCuratedPackage(packageName, name, kubeconfig.FromClusterName(test.ClusterName), constants.EksaPackagesName)
// 		test.VerifyHelloPackageInstalled(name, withMgmtCluster(test))
// 	}
// }

func cleanupCuratedPackageInstall(name string) func(*framework.ClusterE2ETest) {
	return func(test *framework.ClusterE2ETest) {
		alert(test.T, "STARTING cleanupCuratedPackageInstall")
		test.UninstallCuratedPackage(name)
		alert(test.T, "FINISHED cleanupCuratedPackageInstall")
	}
}
