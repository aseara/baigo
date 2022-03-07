package main

import (
	"fmt"
	"time"

	"github.com/aseara/baigo/lesson35/workerpool"
)

func main() {
	p := workerpool.New(5, workerpool.WithBlock(true))

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
