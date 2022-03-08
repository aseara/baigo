package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/aseara/baigo/lesson36/frame"
	"github.com/aseara/baigo/lesson36/packet"
	"github.com/lucasepe/codename"
)

func main() {
	var wg sync.WaitGroup
	num := 5

	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			startClient(i)
		}(i)
	}

	wg.Wait()
}

func startClient(i int) {
	quit := make(chan struct{})
	done := make(chan struct{})

	conn, err := net.Dial("tcp", "172.30.251.20:8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	rbuf := bufio.NewReader(conn)
	wbuf := bufio.NewWriter(conn)
	defer func() { _ = conn.Close() }()
	fmt.Printf("[client %d]: dial ok\n", i)

	frameCodec := frame.NewMyFrameCodec()

	go func() {
		// handle ack
		for {
			select {
			case <-quit:
				done <- struct{}{}
				return
			default:
			}

			_ = conn.SetReadDeadline(time.Now().Add(time.Second * 1))
			ackFramePayload, err := frameCodec.Decode(rbuf)
			if err != nil {
				if e, ok := err.(net.Error); ok {
					if e.Temporary() {
						continue
					}
				}
				panic(fmt.Sprintf("[client %d](:%s): %v", i, conn.LocalAddr().String(), err))
			}
			p, err := packet.Decode(ackFramePayload)
			if err != nil {
				panic(err)
			}
			submitAck, ok := p.(*packet.SubmitAck)
			if !ok {
				panic("not submitack")
			}
			fmt.Printf("[client %d]: the result of submit ack[%s] is %d\n", i, submitAck.ID, submitAck.Result)
			packet.SubmitAckPool.Put(submitAck)
		}
	}()

	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}
	var counter int

	for {
		// send submit
		counter++
		id := fmt.Sprintf("%08d", counter) // 8byte string
		payload := codename.Generate(rng, 4)

		s := packet.SubmitPool.Get().(*packet.Submit)
		s.ID = id
		s.Payload = []byte(payload)

		framePayload, err := packet.Encode(s)

		packet.SubmitPool.Put(s)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[client %d]: send submit id = %s, payload=%s, frame length = %d\n",
			i, s.ID, s.Payload, len(framePayload)+4)

		err = frameCodec.Encode(wbuf, framePayload)
		if err != nil {
			panic(err)
		}
		err = wbuf.Flush()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond * 100)
		if counter >= 10 {
			quit <- struct{}{}
			<-done
			fmt.Printf("[client %d]: exit ok\n", i)
			return
		}
	}
}
