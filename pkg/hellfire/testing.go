package hellfire

import (
	"context"
	"log"
	"testing"
	"time"
)

type Load struct {
	name string
}

func Run(s Scenario, t *testing.T) bool {
	t.Log(s.Name)
	ctx, cancel := context.WithCancel(context.Background())
	go func(t *testing.T, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				log.Println("Hello Ishan")
			}
		}
	}(t, ctx)
	time.Sleep(5 * time.Second)
	cancel()
	return true
}
