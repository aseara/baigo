package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	k1 := "test"

	m.Store(k1, 1)
	v, ok := m.Load(k1)
	fmt.Printf("v: %v, ok: %v\n", v, ok)
}
