package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	sip := net.ParseIP("127.0.0.1")

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 12000}

	conn, _ := net.DialUDP("udp", srcAddr, dstAddr)
	defer func() { _ = conn.Close() }()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input lowercase sentence:")
	text, _ := reader.ReadString('\n')

	_, _ = conn.Write([]byte(text))

	data := make([]byte, 2048)
	n, _ := conn.Read(data)

	fmt.Printf("Receive msg: %s\n", data[:n])
}
