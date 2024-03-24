package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	i := 0
	var mu sync.Mutex

	wg.Add(1)
	go func(mu1 *sync.Mutex) {
		mu1.Lock()
		i = 10
		time.Sleep(10 * time.Second)
		fmt.Printf("g1: i = %d\n", i)
		mu1.Unlock()
		wg.Done()
	}(&mu)

	time.Sleep(time.Second)

	mu.Lock()
	i = 1
	fmt.Printf("main: i = %d\n", i)
	mu.Unlock()

	wg.Wait()
}
