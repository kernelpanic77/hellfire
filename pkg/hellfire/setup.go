package hellfire

import (
	"testing"

	common "github.com/kernelpanic77/hellfire/common"
	exec "github.com/kernelpanic77/hellfire/internal/executor"
)

var executor *exec.Executor

// should start a seperate goroutine for the test with a different context so that we can kill it 
// somehow initialize the metrics machine which would gather the results of that particular test

func Run(s []common.Scenario, thresholds []common.Threshold, t *testing.T, fn common.Task) (int, error) {
	test_metadata := &common.TestMetadata{Scenarios: s, Thresholds: thresholds, T: t, Iteration: fn}
	executor.RunTest(test_metadata)
	// //fmt.Println(executor)
	return 0, nil
}

// initializes the testing framework for instance environment variables, threads etc.
func Main(m *testing.M) int {
	executor = exec.NewExecutor(m)
	exitcode := executor.Setup()
	return exitcode
}
