package executor

import (
	"context"
	"fmt"

	// "fmt"
	"log"
	"testing"

	"github.com/kernelpanic77/hellfire/common"
	"github.com/kernelpanic77/hellfire/internal/metrics"
)

// This is the brain of hellfire
// manages all the tests ->
var logger *log.Logger

type Executor struct {
	Main *testing.M
	exec_ctx context.Context
	tests []*TestManager
}

func NewExecutor(m *testing.M) *Executor {
	tests := make([]*TestManager, 0)
	return &Executor{Main: m, exec_ctx: context.Background(), tests: tests}
}

// start setup of processes that are more time consuming 
func (e *Executor) Setup() int {
	metrics.RegistryOfRegistry = make(metrics.RegistryMap)
	exitcode := e.Main.Run()
	return exitcode
}

func (e *Executor) RunTest(testmetadata *common.TestMetadata) {
	tm := NewTestManager(e.exec_ctx, testmetadata)
	// this would not start running the tests
	e.tests = append(e.tests, tm)
	tm.init()
	// // and initialize all the goroutines on the test
	tm.setup()
	// // will pass on control to the manager
	tm.start()
	tm.waitgroup.Wait()
	
	// @ishanwar: TODO: Should only kill the ingester once all the metrics channel is empty and the Buffer of the ingester is also empty 

	// once we have a confirmed completion, trigger the metrics table prep 
	// //fmt.Println("Registry", tm.report.Registry)
	report_for_term := tm.report.GenerateReport()
	fmt.Println(report_for_term)
}
