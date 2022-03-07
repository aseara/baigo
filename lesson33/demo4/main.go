package main

import (
	"log"
	"sync"
	"time"
)

var (
	_active = make(chan struct{}, 3)
	_job    = make(chan int, 10)
)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			_job <- (i + 1)
		}
		close(_job)
	}()

	var wg sync.WaitGroup

	for j := range _job {
		wg.Add(1)
		go func(j int) {
			_active <- struct{}{}
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-_active
			wg.Done()
		}(j)
	}
	wg.Wait()
}
