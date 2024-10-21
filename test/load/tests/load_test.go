package tests

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/kernelpanic77/hellfire/pkg/hellfire"
	"github.com/kernelpanic77/hellfire/pkg/hellfire/http"
	// "google.golang.org/protobuf/internal/flags"
)

func iteration(t *hellfire.T, client *http.Client) bool {
	res, err := client.Request("GET", "https://httpbin.test.k6.io/")
	log.Println(res)
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
			return len(body) == 9593
		},
	}, "check_list")
	return true
}

func TestFireCreateListLoad(t *testing.T) {
	scenarios := []hellfire.Scenario{
		{
			Name:            "warmup",
			PreAllocatedVUs: 10,
			Stages: []hellfire.Stage{
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
	}
	thresholds := []hellfire.Threshold{
		{
			ScenarioName: "warmup",
			MetricName:   "http_req_duration",
			Conditions:   []string{"p(90) < 400", "p(95) < 800", "p(99.9) < 2000"},
		},
	}
	hellfire.Run(scenarios, thresholds, t, iteration)
}
