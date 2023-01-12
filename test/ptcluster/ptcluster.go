package ptcluster

import (
	"testing"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/test/framework"
)

// Manager of a persistent test cluster.
type Manager interface {
	Setup() error
	Teardown()
	WithCluster(*testing.T, Matcher, TestFunc, CleanupFunc)
	SkipTeardownOnError() bool
}

// Matcher identifies if a given cluster can support an end-to-end test.
type Matcher interface {
	Matches(cluster.Spec) bool
}

// Builder for persistent test clusters.
type Builder func(*testing.T, framework.Provider) *framework.ClusterE2ETest

// TestFunc triggers a test execution.
type TestFunc func(*framework.ClusterE2ETest)

// CleanupFunc deletes a cluster after all tests have completed.
type CleanupFunc func(*framework.ClusterE2ETest)
