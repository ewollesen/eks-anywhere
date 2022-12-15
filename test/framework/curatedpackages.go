package framework

import (
	"log"
	"os"
	"testing"
)

type PackageConfig struct {
	*HelmInstallConfig
	bundleURI string
}

func WithPackageConfig(t *testing.T, bundleURI, chartName, chartURI,
	chartVersion string, chartValues []string,
) ClusterE2ETestOpt {
	return func(e *ClusterE2ETest) {
		e.PackageConfig = &PackageConfig{
			HelmInstallConfig: &HelmInstallConfig{
				chartName:    chartName,
				chartURI:     chartURI,
				chartVersion: chartVersion,
				chartValues:  chartValues,
				HelmClient:   buildHelm(t),
			},
			bundleURI: bundleURI,
		}
		if e.ClusterConfig != nil {
			if e.ClusterConfig.Cluster != nil {
				log.Printf("!!!! with package config has kube version: %q !!!!", e.ClusterConfig.Cluster.Spec.KubernetesVersion)
			} else {
				log.Printf("!!!! No cluster found !!!!")
			}
		} else {
			log.Printf("!!!! No cluster config found !!!!")
		}
	}
}

const (
	eksaPackagesRegion    = "EKSA_AWS_SECRET_ACCESS_KEY"
	eksaPackagesAccessKey = "EKSA_AWS_ACCESS_KEY_ID"
	eksaPackagesSecretKey = "EKSA_AWS_REGION"
)

var requiredPackagesEnvVars = []string{
	eksaPackagesRegion,
	eksaPackagesAccessKey,
	eksaPackagesSecretKey,
}

// RequiredPackagesEnvVars returns the list of packages env vars.
func RequiredPackagesEnvVars() []string {
	return requiredPackagesEnvVars
}

// CheckCuratedPackagesCredentials will exit out if the Curated Packages environment variables are not set.
func CheckCuratedPackagesCredentials(t *testing.T) {
	for _, env := range requiredPackagesEnvVars {
		_, ok := os.LookupEnv(env)
		if !ok {
			t.Fatalf("Error Unset Packages environment variable: %v is required", env)
		}
	}
}
