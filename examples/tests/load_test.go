package tests

import (
	"io"
	"testing"

	"github.com/kernelpanic77/hellfire/common"
	"github.com/kernelpanic77/hellfire/pkg/hellfire"
)

func iteration(t common.Test, client common.Client) bool {
	res, err := client.Request("GET", "https://httpbin.test.k6.io/")
	if err != nil {
		panic(err)
	}
	t.Check(res, common.CheckFuncMap{
		"status_code_check": func(interface{}) bool {
			return res.StatusCode == 200
		},
		"body_size_check": func(interface{}) bool {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return false
			}
			return len(body) == 9593
		},
	}, "check_list")
	return true
}

func TestFireCreateListLoad(t *testing.T) {
	scenarios := []common.Scenario{
		{
			Strategy: "rampup",
			Name:            "warmup",
			PreAllocatedVUs: 10,
			Stages: []common.Stage{
				{
					Target:   5,
					Duration: 0,
				},
				{
					Target:   5,
					Duration: 2,
				},
				{
					Target:   15,
					Duration: 10,
				},
			},
		},
		{
			Strategy: "shared-iterations",  
			Name: "warmup", 
			PreAllocatedVUs: 10,
			Iterations: 10,  
			MaxDuration: 20,
		},
	}
	thresholds := []common.Threshold{
		{
			ScenarioName: "warmup",
			MetricName:   "http_req_duration",
			Conditions:   []string{"p(90) < 400", "p(95) < 800", "p(99.9) < 2000"},
		},
	}
	hellfire.Run(scenarios, thresholds, t, iteration)
}

