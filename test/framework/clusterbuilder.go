package framework

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
)

// Build clusters to use in end-to-end tests.
//
// Aimed at building clusters that are not the system under test, but rather
// are a requirement for exercising the system under test.

// ClusterForE2E is a terrible name. Come up with something better. Maybe SupportCluster?
type ClusterForE2E struct {
	Config   *cluster.Config
	Provider Provider
}

// NewClusterBuilder TODO fill me in.
func NewClusterBuilder(config *cluster.Config, provider Provider) *ClusterForE2E {
	return &ClusterForE2E{
		Config:   config,
		Provider: provider,
	}
}

// Build TODO fill me in.
func (c *ClusterForE2E) Build(ctx context.Context) error {
	// We're gonna need some access to an EKSA binary...
	return nil
}

// Teardown TODO fill me in.
func (c *ClusterForE2E) Teardown(ctx context.Context) error { return nil }

// // runEKSA TODO fill me in.
// func (c *ClusterForE2E) runEKSA(ctx context.Context, args ...any) error {
// 	return nil
// }

// func (e *ClusterE2ETest) Run(name string, args ...string) {
// 	command := strings.Join(append([]string{name}, args...), " ")
// 	shArgs := []string{"-c", command}

// 	// This log message can come after e.T has finished, and that causes a
// 	// panic. What can be done about that? Is it as simple as updating e.T to
// 	// point to test2? Or even another test entirely?
// 	e.T.Log("Running shell command", "[", command, "]")
// 	cmd := exec.CommandContext(context.Background(), "sh", shArgs...)

// 	envPath := os.Getenv("PATH")

// 	workDir, err := os.Getwd()
// 	if err != nil {
// 		e.T.Fatalf("Error finding current directory: %v", err)
// 	}

// 	var stdoutAndErr bytes.Buffer

// 	cmd.Env = os.Environ()
// 	cmd.Env = append(cmd.Env, fmt.Sprintf("PATH=%s/bin:%s", workDir, envPath))
// 	cmd.Stderr = io.MultiWriter(os.Stderr, &stdoutAndErr)
// 	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutAndErr)

// 	if err = cmd.Run(); err != nil {
// 		scanner := bufio.NewScanner(&stdoutAndErr)
// 		var errorMessage string
// 		// Look for the last line of the out put that starts with 'Error:'
// 		for scanner.Scan() {
// 			line := scanner.Text()
// 			if strings.HasPrefix(line, "Error:") {
// 				errorMessage = line
// 			}
// 		}

// 		if err := scanner.Err(); err != nil {
// 			e.T.Fatalf("Failed reading command output looking for error message: %v", err)
// 		}

// 		if errorMessage != "" {
// 			if e.ExpectFailure {
// 				e.T.Logf("This error was expected. Continuing...")
// 				return
// 			}
// 			e.T.Fatalf("Command %s %v failed with error: %v: %s", name, args, err, errorMessage)
// 		}

// 		e.T.Fatalf("Error running command %s %v: %v", name, args, err)
// 	}
// }
