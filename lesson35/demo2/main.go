package main

import (
	"fmt"
	"time"

	"github.com/aseara/baigo/lesson35/workerpool"
)

func main() {
	p := workerpool.New(5, workerpool.WithBlock(false), workerpool.WithPreAllocWorkers(false))

	time.Sleep(time.Second * 2)
	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
		})
		if err != nil {
			fmt.Printf("task: %d, err: %v\n", i, err)
		}
	}

	p.Free()
}
