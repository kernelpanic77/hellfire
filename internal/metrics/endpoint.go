package metrics

import (
	"context"
	"sync"
)

type Endpoint interface {
	AddSamples(metrics []SampleContainer) // allows to add Samples to the Endpoint Buffer
	Start(wg *sync.WaitGroup, ctx context.Context)      // allows to start a goroutine which pumps samples from the buffers
}
