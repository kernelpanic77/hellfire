package hellfire

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/schollz/progressbar/v3"
)

var loc *string = flag.String("location", "World", "Name of location to greet")

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
func Run(s Scenario, t *testing.T, fn task) {
	// somehow initialize the metrics machine which would gather the results of that particular test
	a := &Artillary{init_workers: 0}
	// defer a.stop()
	a.start(&s)
	done := make(chan bool)
	a.Fire(fn, done)
	<-done
}

// initializes the testing framework for instance environment variables, threads etc.
func Main(m *testing.M) int {
	// fmt.Println(loc)
	m.Run()
	return 0
}
