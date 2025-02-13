package metrics

import (
	"context"
	"sync"
	"time"
)

var (
	_ Endpoint = &Ingester{}
)

const duration = 50 * time.Millisecond

type SampleBuffer struct {
	mu     *sync.Mutex
	buffer []SampleContainer
	maxLen int
}

func (sb *SampleBuffer) PushSamples(samples []SampleContainer) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	//fmt.Println("PushSamples")
	//fmt.Println(samples)
	sb.buffer = append(sb.buffer, samples...)
	//fmt.Println(sb.buffer)
}

func (sb *SampleBuffer) FetchSamples() []SampleContainer {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	samples, len_samples := sb.buffer, len(sb.buffer)
	if len_samples > sb.maxLen {
		sb.maxLen = len_samples
	}
	sb.buffer = make([]SampleContainer, 0, (sb.maxLen+len_samples)/2)
	return samples
}

type PeriodicFlusher struct {
	ticker    *time.Ticker
	buffer  *SampleBuffer
}

func (p *PeriodicFlusher) Flush(ctx context.Context) {
	// should flush metrics to the sink in every interval of time, should be less than a second to keep things accurate
	for {
		select {
		case <-ctx.Done():
			//fmt.Println("Stop Flushing Metrics")
			return
		case <-p.ticker.C:
			p.buffer.send_to_sink()
			// //fmt.Println(p.buffer)
		}
	}
}

type Ingester struct {
	// create a buffer to store the metrics
	buffer  *SampleBuffer
	flusher *PeriodicFlusher
}

func NewIngester() *Ingester {
	buffer := &SampleBuffer{
		mu: &sync.Mutex{},
		buffer: make([]SampleContainer, 0),
		maxLen: 1000,
	}
	return &Ingester{
		buffer:  buffer,
		flusher: &PeriodicFlusher{
			ticker: time.NewTicker(duration),
			buffer: buffer ,
		},
	}
}

func (i *Ingester) Start(wg *sync.WaitGroup, ctx context.Context) { // Should basically start a go routine to flush metrics to respective Metrics Sink
	defer wg.Done()
	i.flusher.Flush(ctx)
	//fmt.Println("golang")
}

func (i *Ingester) AddSamples(samples []SampleContainer) { // Should push the samples to the buffer
	i.buffer.PushSamples(samples)
}

func (b *SampleBuffer) send_to_sink() {
	// read from buffer
	sample_containers := b.FetchSamples()
	for _, container := range sample_containers {
		samples := container.GetSamples()
		for _, sample := range samples {
			// push to the sink of the sample
			mlist := sample.metric
			for _, m := range mlist {
				//fmt.Println(m.Name)
				m.Sink.AddSample(sample)
			}
			// //fmt.Println(m.Sink.FetchSampleValue())
		}
	}
}
