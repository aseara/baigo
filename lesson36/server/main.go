package main

import (
	"fmt"
	"net"

	"github.com/aseara/baigo/lesson36/frame"
	"github.com/aseara/baigo/lesson36/packet"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}
	fmt.Println("server start ok(on *.8888)")
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			break
		}
		// start a new goroutine to handle the new connection.
		go handleConn(c)
	}

}

func handleConn(c net.Conn) {
	defer func() { _ = c.Close() }()
	frameCodec := frame.NewMyFrameCodec()

	for {
		// decode the frame to get the payload
		framePayload, err := frameCodec.Decode(c)
		if err != nil {
			fmt.Println("handleConn: frame decode error: ", err)
			return
		}

		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			fmt.Println("handleConn: handle packet error: ", err)
			return
		}
		err = frameCodec.Encode(c, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error: ", err)
		}
	}
}

func handlePacket(framePayload frame.Payload) (ackFramePayload frame.Payload, err error) {
	var p packet.Packet
	p, err = packet.Decode(framePayload)
	if err != nil {
		fmt.Println("handleConn: packet decode error: ", err)
		return
	}
	switch p := p.(type) {
	case *packet.Submit:
		fmt.Printf("recv submit: id = %s, payload=%s\n", p.ID, string(p.Payload))
		submitAck := &packet.SubmitAck{
			ID:     p.ID,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(submitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode error: ", err)
			return nil, err
		}
		return ackFramePayload, nil
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}
