package metrics

import "context"

type Endpoint interface {
	AddSamples(metrics []SampleContainer) // allows to add Samples to the Endpoint Buffer
	Start(ctx context.Context)      // allows to start a goroutine which pumps samples from the buffers
}
