package metrics

import (
	"context"
	"sync"
	"time"
)

// responsible for reading the metrics chan and sending the data to endpoints

const collectRate = 50 * time.Millisecond

type Broadcaster struct {
	endpoints []Endpoint
	ctx context.Context
}

func NewBroadcaster(endpoints []Endpoint, ctx context.Context) *Broadcaster {
	return &Broadcaster{
		endpoints: endpoints,
		ctx: ctx,
	}
}

func (b *Broadcaster) Start(wg *sync.WaitGroup, containers chan SampleContainer) {
	ticker := time.NewTicker(collectRate)
	for {
		select {
		case <-ticker.C:
			// fmt.Println("Ishan is a fucking genius")
			data := FetchBufferedSamples(containers)
			for _, endpoint := range b.endpoints {
				// this method should be non blocking, because if we are running multiple tests 
				// all of the TestManagers are writing to the buffer we dont want write to be blocked for that endpoint
				// fmt.Println("endpoints")
				// fmt.Println(endpoint)
				for _, samples := range data {
					endpoint.AddSamples([]SampleContainer{samples}) 
				}
			}
		case <-b.ctx.Done(): 
			wg.Done()
			return 
		}
	}
}