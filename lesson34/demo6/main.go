package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a int32 = 13
	addValue := atomic.AddInt32(&a, 1)
	fmt.Println("after add 1, a is", addValue)
	delVale := atomic.AddInt32(&a, -1)
	fmt.Println("after del 1, a is", delVale)
}
