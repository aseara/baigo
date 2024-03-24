package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("task running...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("task is done.")
	}()
	fmt.Println("waiting for task done...")
	time.Sleep(1 * time.Second)
	fmt.Println("end of all.")
}
