package tests

import (
	"io/ioutil"
	"testing"

	"github.com/kernelpanic77/hellfire/pkg/hellfire"
	"github.com/kernelpanic77/hellfire/pkg/hellfire/http"
	// "google.golang.org/protobuf/internal/flags"
)

func iteration(t *hellfire.T, client *http.Client) bool {

	res, err := client.Request("GET", "random URL")
	if err != nil {
		panic(err)
	}
	t.Check(res, hellfire.CheckFuncMap{
		"status_code_check": func(interface{}) bool {
			return res.StatusCode == 200
		},
		"body_size_check": func(interface{}) bool {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return false
			}
			return len(body) == 2593
		},
	}, "check_list")
	return true
}

func TestFireCreateListLoad(t *testing.T) {
	// initialize the test
	// let's call it hellfire init

	scenarios := []hellfire.Scenario{
		{Name: "httpbin_test", Spawn_rate: 1, Target: 5, Duration: 10},
	}

	// planner := hellfire.Planner{
	// 	stages: [],
	// 	thresholds: []
	// }

	// Define the Request URL, headers, Body etc.
	// Start the process
	for _, tc := range scenarios {
		t.Logf("Starting scenario, %s", tc.Name)
		hellfire.Run(tc, t, iteration)
	}

	// hellfire finish callbacks
}
