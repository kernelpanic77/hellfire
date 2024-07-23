package main

import (
	"fmt"
	"time"

	"github.com/panjf2000/ants"
)

func main() {
	// a, err := ants.NewPool(1, ants.WithNonblocking(true))

	done := make(chan bool)
	task := func(i interface{}) {
		ticker := time.NewTicker(1 * time.Second)
		timer := time.NewTimer(10000 * time.Second)
		for range ticker.C {
			// case <-ticker.C:
			// 	fmt.Println("Done!")
			// default:
			select {
			case <-timer.C:
				ticker.Stop()
				done <- true
				return
			default:
				// fmt.Println(i.(int))
			}
		}
		// done <- true
		// done <- true
	}
	// fmt.Print(a.Running())

	p, _ := ants.NewPoolWithFunc(1000, func(i interface{}) {
		task(i)
	})
	defer p.Release()
	for i := 0; i < 1001; i++ {
		p.Invoke(100)
		fmt.Println(p.Running())
	}

	<-done
}
