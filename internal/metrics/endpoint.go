package metrics

type Endpoint interface {
	AddSamples() // allows to add Samples to the Endpoint Buffer
	Start()      // allows to start a goroutine which pumps samples from the buffers
}
