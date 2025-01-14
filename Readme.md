# Hellfire: A Go-based Load Testing Package

**Hellfire** is a lightweight, customizable load testing library built specifically for Go developers. Unlike traditional load testing tools such as [k6](https://k6.io/), [Locust](https://locust.io/), and [Gatling](https://gatling.io/), Hellfire leverages the built-in testing framework provided by Go's default packages. This makes it highly extensible and integrated into the Go ecosystem, allowing developers to perform load and spike testing directly within the familiar context of Go's `testing` package.

## Key Features:
- **Built on Go's Standard Testing Framework**: Hellfire uses Go's native testing tools, making it easy for Go developers to integrate load testing into their existing test suites.
- **Highly Customizable**: Unlike other tools, Hellfire enables you to write and modify tests with full control, utilizing Go's powerful language features and libraries.
- **Seamless Integration**: No need for external dependencies—just leverage Go’s default packages to simulate real-world load, spikes, and stress tests.
- **Ideal for Go-Only Environments**: Hellfire is tailored for developers who prefer staying within the Go ecosystem, avoiding the need for managing separate load testing tools or environments.

## Why Hellfire?

### Traditional Tools vs. Hellfire:
While tools like k6, Locust, and Gatling provide feature-rich environments for load testing, they often come with a learning curve and require developers to step outside the Go ecosystem. Hellfire, on the other hand, is designed to work seamlessly within Go's standard testing framework, providing:

- **Easier Setup**: No external installations or dependencies are required. Simply write Go tests, and you’re ready to perform load testing.
- **Full Customization**: Hellfire takes full advantage of Go's testing capabilities, allowing you to tweak and extend tests as needed to match your specific requirements.
- **Built-In Concurrency**: With Go’s goroutines and channels, you can simulate complex concurrent load scenarios effortlessly.

Basic Example: 

```go 
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
```