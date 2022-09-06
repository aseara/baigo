package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 12000}

	listener, _ := net.ListenUDP("udp", srcAddr)

	fmt.Printf("Local <%s> \n", listener.LocalAddr().String())

	data := make([]byte, 2048)

	for {
		n, rAddr, _ := listener.ReadFromUDP(data)

		fmt.Printf("<%s> %s\n", rAddr, data[:n])

		up := strings.ToUpper(string(data[:n]))

		_, _ = listener.WriteToUDP([]byte(up), rAddr)
	}
}
