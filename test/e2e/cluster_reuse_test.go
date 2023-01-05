//go:build e2e
// +build e2e

package e2e

import (
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
	m.e2eTest = nil
	return nil
}

func (m *vsphereClusterManager) Teardown() {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Tests are completed at this point, so any logging will cause a panic
	m.e2eTest.T = &testing.T{}
	m.e2eTest.StopIfFailed()
	m.e2eTest.DeleteCluster()
	m.e2eTest = nil
}

type runFunc func(*framework.ClusterE2ETest)

type cleanupFunc func(*framework.ClusterE2ETest)

func (m *vsphereClusterManager) WithCluster(t *testing.T, run runFunc, cleanup cleanupFunc) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.e2eTest == nil {
		test := framework.NewClusterE2ETest(t,
			framework.NewVSphere(t, framework.WithBottleRocket123()),
			framework.WithClusterFiller(api.WithKubernetesVersion(v1alpha1.Kube123)),
			framework.WithPackageConfig(t, packageBundleURI(v1alpha1.Kube121),
				EksaPackageControllerHelmChartName, EksaPackageControllerHelmURI,
				EksaPackageControllerHelmVersion, EksaPackageControllerHelmValues))
		test.GenerateClusterConfig()
		test.CreateCluster()
		m.e2eTest = test
	} else {
		t.Logf("Re-using existing test cluster %q", m.e2eTest.ClusterName)
		m.e2eTest.T = t
	}

	defer func() {
		m.e2eTest.T = framework.NewLoggingOnlyT()
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
	os.Exit(code)
}

func TestClusterReuse(s *testing.T) {
	s.Run("test number one", func(t *testing.T) {
		manager.WithCluster(s, runCuratedPackageInstallWithName("test1"), cleanupCuratedPackageInstall("test1"))
	})

	s.Run("test number two", func(t *testing.T) {
		manager.WithCluster(s, runCuratedPackageInstallWithName("test2"), cleanupCuratedPackageInstall("test2"))
	})
}

func cleanupCuratedPackageInstall(name string) func(*framework.ClusterE2ETest) {
	return func(test *framework.ClusterE2ETest) {
		test.UninstallCuratedPackage(name)
	}
}
