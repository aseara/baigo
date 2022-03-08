package main

import (
	"bufio"
	"fmt"
	"net"

	"net/http"
	_ "net/http/pprof"

	"github.com/aseara/baigo/lesson36/frame"
	"github.com/aseara/baigo/lesson36/packet"

	"github.com/aseara/baigo/lesson36/metrics"
)

func main() {
	go func() {
		_ = http.ListenAndServe(":6060", nil)
	}()
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
	metrics.ClientConnected.Inc()
	defer func() {
		_ = c.Close()
		metrics.ClientConnected.Dec()
	}()
	frameCodec := frame.NewMyFrameCodec()
	rbuf := bufio.NewReader(c)
	wbuf := bufio.NewWriter(c)
	for {
		// decode the frame to get the payload
		framePayload, err := frameCodec.Decode(rbuf)
		if err != nil {
			fmt.Println("handleConn: frame decode error: ", err)
			return
		}
		metrics.ReqRecvTotal.Add(1)

		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			fmt.Println("handleConn: handle packet error: ", err)
			return
		}
		err = frameCodec.Encode(wbuf, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error: ", err)
		}
		err = wbuf.Flush()
		if err != nil {
			fmt.Println("handleConn: write flash error: ", err)
		}
		metrics.RspSendTotal.Add(1)
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
		submitAck := packet.SubmitAckPool.Get().(*packet.SubmitAck)
		submitAck.ID = p.ID
		submitAck.Result = 0
		packet.SubmitPool.Put(p)
		ackFramePayload, err = packet.Encode(submitAck)
		packet.SubmitAckPool.Put(submitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode error: ", err)
			return nil, err
		}
		return ackFramePayload, nil
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}
