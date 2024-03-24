package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		ii := i
		go func() {
			defer wg.Done()
			fmt.Printf("%d finished\n", ii)
		}()
	}

	wg.Wait()
	fmt.Println("all finished")
}
