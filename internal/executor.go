package internal

import (
	"log"
	"os"
	"testing"

	"github.com/kernelpanic77/hellfire/pkg/hellfire"
)

// This is the brain of hellfire manages all the tests

var logger *log.Logger

type Executor struct {
	Main *testing.M
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
}

func (e *Executor) setupArtillary() {
	// responsible for activating the initial set of workers
	//  as well as policy of spawning of workers
}

func (e *Executor) setupExternalEndpoints() {
	// setup all the backends which are needed apart from ingester (like TimeScaleDB)
}

func (e *Executor) setupThresholdObservers() {
	// great idea if your tests needs to observe the thresholds constantly for some reason as part of some metric
}

func (e *Executor) setupFundamentalMetrics() {

}

func (e *Executor) RunTest(t *hellfire.TestMetadata) {
	// this would not start running the tests

	// and initialize all the goroutines on the test

	// will pass on control to the manager
}
