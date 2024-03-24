package main

import (
	"fmt"

	"github.com/aseara/baigo/lesson36/packet"
)

func main() {
	s := packet.Submit{
		ID: "hello",
	}

	addr := &s.ID

	fmt.Printf("s.ID: %s", *addr)
}
