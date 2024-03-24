package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
		ch2 <- 7
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch1:
			if ok {
				fmt.Println(x)
				continue
			}
			ch1 = nil

		case x, ok := <-ch2:
			if ok {
				fmt.Println(x)
				continue
			}
			ch2 = nil
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("program end")
}
