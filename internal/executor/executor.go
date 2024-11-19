package internal

import (
	artillary "github.com/kernelpanic77/hellfire/internal/artillary"
	state "github.com/kernelpanic77/hellfire/internal/state"

	"log"
	"os"
	"testing"

	metrics "github.com/kernelpanic77/hellfire/internal/metrics"
)

// This is the brain of hellfire manages all the tests

var logger *log.Logger

type Executor struct {
	Main *testing.M
	logger *log.Logger 
	State *state.TestState
	registry *metrics.MetricsRegistry
	artillary *artillary.Artillary
	endpoints []metrics.Endpoint
	sinks map[metrics.MetricType]metrics.Sink
}

type TestMetadata interface {
	Validate() bool
}

func NewExecutor(m *testing.M) *Executor {
	return &Executor{Main: m}
}

func (e *Executor) setMain(m *testing.M) {
	e.Main = m
}

// start tests
func (e *Executor) start() int {
	exitcode := e.Main.Run()
	return exitcode
}

// initialize all the variables and config (setup)
func (e *Executor) Setup(m *testing.M) (int, error) {
	logger = log.New(os.Stdout, "INFO: HELLFIRE", log.LstdFlags)
	logger.Println("starting setup")
	e.setMain(m)
	return 0, nil
}

func (e *Executor) setupMetricsEngine() {
	// resposible for collecting metrics
	e.registry = &metrics.MetricsRegistry{}
}

func (e *Executor) setupArtillary() {
	// responsible for activating the initial set of workers
	//  as well as policy of spawning of workers
	e.artillary = &artillary.Artillary{}
}

func (e *Executor) setupExternalEndpoints() {
	ingestor := &metrics.Ingester{}
	// setup all the backends which are needed apart from ingester (like TimeScaleDB)
	e.endpoints = append(e.endpoints, ingestor)
}

func (e *Executor) setupThresholdObservers() {
	// great idea if your tests needs to observe the thresholds constantly for some reason as part of some metric

}


func (e *Executor) setupMetricSinks() {
	// done with setup of basic metrics
	e.sinks[metrics.Counter] =  metrics.NewSink(metrics.Counter)
	e.sinks[metrics.Guage] =  metrics.NewSink(metrics.Guage)
	e.sinks[metrics.Rate] = metrics.NewSink(metrics.Rate) 
	e.sinks[metrics.Trend] = metrics.NewSink(metrics.Trend)
}

// func (e *Executor) setupFundamentalMetrics() {
// 	metrics.InitDefaultMetrics(e.registry, e.sinks)
// 	metrics.InitHTTPMetrics(e.registry, e.sinks)
// }



func (e *Executor) RunTest(t TestMetadata) {
	// this would not start running the tests
	
	// and initialize all the goroutines on the test

	// will pass on control to the manager
}
