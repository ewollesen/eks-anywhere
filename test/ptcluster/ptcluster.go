package ptcluster

import (
	"testing"

	"github.com/aws/eks-anywhere/test/framework"
)

// Manager of a persistent test cluster.
type Manager interface {
	// Setup before any tests are run.
	Setup() error
	// Teardown all clusters still running after all tests are finished.
	Teardown()
	// WithCluster selects a cluster and runs the test on it.
	WithCluster(*testing.T, Requirements, TestFunc, CleanupFunc)
	// SkipTeardownOnError if you prefer the cluster remain operational after
	// a failed test.
	SkipTeardownOnError() bool
}

// Requirements identifies if a given cluster can support an end-to-end test.
type Requirements struct {
	Providers []Provider
}

// Builder for persistent test clusters.
type Builder func(*testing.T, framework.Provider) *framework.ClusterE2ETest

// TestFunc triggers a test execution.
type TestFunc func(*framework.ClusterE2ETest)

// CleanupFunc deletes a cluster after all tests have completed.
type CleanupFunc func(*framework.ClusterE2ETest)

// Provider describes a persistent test cluster's provider.
//
//go:generate stringer -type Provider -trimprefix Provider
type Provider int

const (
	// ProviderAny matches any Provider.
	ProviderAny Provider = iota
	// ProviderVsphere matches Vsphere.
	ProviderVsphere
	// ProviderDocker matches Docker.
	ProviderDocker
	// ProviderSnow matches Snow.
	ProviderSnow
	// ProviderBareMetal matches Bare Metal.
	ProviderBareMetal
	// ProviderCloudStack matches CloudStack.
	ProviderCloudStack
)

// MatchesProvider between a Manager and a test.
func (r Requirements) MatchesProvider(other Provider) bool {
	if len(r.Providers) == 0 {
		return true
	}
	for _, p := range r.Providers {
		if p == ProviderAny || p == other {
			return true
		}
	}
	return false
}
