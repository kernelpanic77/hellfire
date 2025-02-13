package executor

import (
	"context"
	"sync"
	"testing"

	"github.com/kernelpanic77/hellfire/common"
	artillary "github.com/kernelpanic77/hellfire/internal/artillary"
	metrics "github.com/kernelpanic77/hellfire/internal/metrics"
	state "github.com/kernelpanic77/hellfire/internal/state"
)

// Run the TestManagement goroutine
// executor creates a goroutine which manages the test
// this is so that we can kill the goroutine managing the Test at any point if necessary -> will kill all the necessary allocations

type TestManager struct {
	Main *testing.M
	State *state.TestState
	waitgroup *sync.WaitGroup
	machine *metrics.Machine
	infantry *artillary.Infantry
	test_context context.Context
	test_termination context.CancelFunc
	broadcaster *metrics.Broadcaster
	testMetadata *common.TestMetadata
	report *state.Report
}

func NewTestManager(ctx context.Context, testMetadata *common.TestMetadata) *TestManager {
	ctx_term, cancel_term := context.WithCancel(context.Background())
	return &TestManager{
		test_context: ctx_term,
		test_termination: cancel_term,
		testMetadata: testMetadata,
		waitgroup: &sync.WaitGroup{},
	}
}

// Complete setup of the metrics machine 
func (tm *TestManager) init() {
	// setup the metrics machine
	registry := metrics.NewMetricsRegistry()
	report := state.NewReport(tm.test_context, registry)
	// //fmt.Println("Registry")
	// //fmt.Println(tm.report.Registry)
	tm.report = report
	// const registryKey string = "registry"
	metrics.RegistryOfRegistry["test_name"] = registry
	metrics.InitDefaultMetrics(registry)
	metrics.InitHTTPMetrics(registry)
	tm.machine = metrics.NewMachine(tm.test_context, registry)

	// // endpoints
	endpoints := tm.machine.GetEndpoints()
	
	// // broadcaster
	tm.broadcaster = metrics.NewBroadcaster(endpoints, tm.test_context)
	// infantry setup
	tm.infantry = artillary.NewInfantry(tm.test_context, tm.testMetadata, tm.machine.GetSamplesChan())
}

// Complete setup of external endpoints if any needed
// which means start the flush goroutines, which constantly flush data from the endpoint buffers to the Endpoints 
func (tm *TestManager) setup() {
	// start goroutines if needed
	tm.waitgroup.Add(1)
	go tm.broadcaster.Start(tm.waitgroup, tm.machine.GetSamplesChan())
	
	for _, endpoint := range(tm.machine.GetEndpoints()) {
		tm.waitgroup.Add(1)
		go endpoint.Start(tm.waitgroup, tm.test_context)
	}	
}

// should start the goroutine to trigger the infantry 
func (tm *TestManager) start() {
	tm.waitgroup.Add(1)
	//fmt.Println("starting the test manager")
	go tm.infantry.Action(tm.waitgroup, tm.test_termination)
}
