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
	"github.com/aws/eks-anywhere/test/ptcluster"
)

var manager ptcluster.Manager

type vsphereOnlyClusterManager struct {
	mu             sync.Mutex
	test           *framework.ClusterE2ETest
	failed         bool
	clusterBuilder ptcluster.Builder
}

func (m *vsphereOnlyClusterManager) Setup() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.test = nil
	return nil
}

func (m *vsphereOnlyClusterManager) SkipTeardownOnError() bool {
	return false
}

func (m *vsphereOnlyClusterManager) Teardown() {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Tests are completed at this point, so any logging will cause a panic
	t := framework.NewLoggingOnlyT()
	m.test.T = t
	if !m.failed || m.SkipTeardownOnError() {
		m.test.DeleteCluster()
	}
	m.test = nil
}

func (m *vsphereOnlyClusterManager) WithCluster(t *testing.T, rqmts ptcluster.Requirements, run ptcluster.TestFunc, cleanup ptcluster.CleanupFunc) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !rqmts.MatchesProvider(ptcluster.ProviderVsphere) {
		t.Logf("persistent test cluster manager is unable to provide a matching cluster")
		t.SkipNow()
	}

	if m.test == nil {
		test := framework.NewClusterE2ETest(t,
			framework.NewVSphere(t, framework.WithBottleRocket123()),
			framework.WithClusterFiller(api.WithKubernetesVersion(v1alpha1.Kube123)),
			framework.WithPackageConfig(t, packageBundleURI(v1alpha1.Kube121),
				EksaPackageControllerHelmChartName, EksaPackageControllerHelmURI,
				EksaPackageControllerHelmVersion, EksaPackageControllerHelmValues))
		test.GenerateClusterConfig()
		test.CreateCluster()
		m.test = test
	} else {
		t.Logf("Re-using existing test cluster %q", m.test.ClusterName)
		m.test.T = t
	}

	defer func() {
		m.test.T = framework.NewLoggingOnlyT()
		cleanup(m.test)
	}()
	run(m.test)
	m.failed = m.failed || m.test.T.Failed()
}

func init() {
	builder := func(t *testing.T, p framework.Provider) *framework.ClusterE2ETest {
		return framework.NewClusterE2ETest(t, p)
	}
	manager = &vsphereOnlyClusterManager{clusterBuilder: builder}
}

var _ ptcluster.Manager = (*vsphereOnlyClusterManager)(nil)

func TestMain(m *testing.M) {
	if err := manager.Setup(); err != nil {
		log.Fatal(err)
	}
	defer manager.Teardown()
	code := m.Run()
	if code != 0 && manager.SkipTeardownOnError() {
		os.Exit(code)
	}
}

func TestClusterReuse(s *testing.T) {
	rqmts := ptcluster.Requirements{
		Providers: []ptcluster.Provider{ptcluster.ProviderVsphere},
	}

	s.Run("test number one", func(t *testing.T) {
		manager.WithCluster(s, rqmts, runCuratedPackageInstallWithName("test1"), cleanupCuratedPackageInstall("test1"))
	})

	s.Run("test number two", func(t *testing.T) {
		manager.WithCluster(s, rqmts, runCuratedPackageInstallWithName("test2"), cleanupCuratedPackageInstall("test2"))
	})
}

func TestClusterReuseThree(s *testing.T) {
	rqmts := ptcluster.Requirements{
		Providers: []ptcluster.Provider{ptcluster.ProviderVsphere},
	}

	manager.WithCluster(s, rqmts, runCuratedPackageInstallWithName("test3"), cleanupCuratedPackageInstall("test3"))
}

func TestClusterReuseFourIsSkipped(s *testing.T) {
	rqmts := ptcluster.Requirements{
		Providers: []ptcluster.Provider{ptcluster.ProviderDocker},
	}

	manager.WithCluster(s, rqmts, runCuratedPackageInstallWithName("test4"), cleanupCuratedPackageInstall("test4"))
}

func cleanupCuratedPackageInstall(name string) func(*framework.ClusterE2ETest) {
	return func(test *framework.ClusterE2ETest) {
		test.UninstallCuratedPackage(name)
	}
}
