package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		fmt.Println("Hello Defer!")
	}()

	fmt.Println("Hello World!")

	os.Exit(0)
}
