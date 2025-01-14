package artillary

import (
	"context"
	"sync"

	"github.com/kernelpanic77/hellfire/common"
	"github.com/kernelpanic77/hellfire/internal/metrics"
)

// responsible for managing the multiple types of artillary

type Infantry struct {
	ctx context.Context
	Artillaries map[string]*Artillary
	Scenarios []common.Scenario
}

func NewInfantry(ctx context.Context, testMetadata *common.TestMetadata, metrics_chan chan metrics.SampleContainer) *Infantry {
	artillaries := make(map[string]*Artillary)
 	for _, scenario := range testMetadata.Scenarios {
		artillaries[scenario.Name] = NewArtillary(ctx, scenario, testMetadata.Iteration, metrics_chan)
	}
	return &Infantry{ctx: ctx, Artillaries: artillaries, Scenarios: testMetadata.Scenarios }
}

func (infantry *Infantry) Action(wg *sync.WaitGroup, termination context.CancelFunc) {
	for _, scenario := range infantry.Scenarios {
		infantry.triggerScenario(scenario)
	}
	wg.Done()
	termination()
}

func (infantry *Infantry) triggerScenario(scenario common.Scenario) {
	infantry.Artillaries[scenario.Name].StartArtillary()
}