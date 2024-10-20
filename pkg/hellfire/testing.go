package hellfire

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"testing"
	"time"

	exec "github.com/kernelpanic77/hellfire/internal"

	"github.com/schollz/progressbar/v3"
)

var loc *string = flag.String("location", "World", "Name of location to greet")

var executor *exec.Executor

type Load struct {
	name string
}

// This struct should define functions like LogF, FatalF, ErrorF etc.
type T struct {
	artillary *Artillary
}

func (t *T) ProgressBar() {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
}

func (t *T) Log(s string) {
	log.Println(s)
}

type CheckFunc func(interface{}) bool

type CheckFuncMap map[string]CheckFunc

func (t *T) Check(val interface{}, checks CheckFuncMap, tag string) bool {
	check_val := true

	for check_name, check_func := range checks {
		curr_check := check_func(val)
		if !curr_check {
			// t.("%s has failed!", check_name)
			t.Log(fmt.Sprintf("%s check has failed!", check_name))
		}
		check_val = check_val && curr_check
	}

	return check_val
}

func (t *T) Fatal(msg string) {
	t.Log(msg)
	panic(errors.New(msg))
}

// Runs each scenario, where s defines the scenario and fn describes the checks to be performed in each iteration
// func Run(s []Scenario, thresholds []Threshold, t *testing.T, fn task) {
// 	// somehow initialize the metrics machine which would gather the results of that particular test
// 	a := &Artillary{init_workers: 0}
// 	// defer a.stop()
// 	a.start(&s)
// 	done := make(chan bool)
// 	a.Fire(fn, done)
// 	<-done
// }

func Run(s []Scenario, thresholds []Threshold, t *testing.T, fn task) (int, error) {
	// somehow initialize the metrics machine which would gather the results of that particular test
	task_metadata := &TestMetadata{Scenarios: s, Thresholds: thresholds, T: t, Iteration: fn}
	executor.RunTest(task_metadata)
	return 0, nil
}

// initializes the testing framework for instance environment variables, threads etc.
func Main(m *testing.M) int {
	exitcode, err := executor.Setup(m)
	if err != nil {
		panic(err)
	}
	return exitcode
}
