package frame

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// Payload payload
type Payload []byte

// StreamFrameCodec frame codec
type StreamFrameCodec interface {
	Encode(io.Writer, Payload) error   // data -> frame, 并写入io.Writer
	Decode(io.Reader) (Payload, error) // 从 io.Reader 中提取 frame payload, 并返回给上层
}

var (
	// ErrShortWrite short write error
	ErrShortWrite = errors.New("short write")
	// ErrShortRead short read error
	ErrShortRead = errors.New("short read")
)

type myFrameCodec struct{}

// NewMyFrameCodec factory method for myFrameCodec
func NewMyFrameCodec() StreamFrameCodec {
	return &myFrameCodec{}
}

func (p *myFrameCodec) Encode(w io.Writer, f Payload) error {
	totalLen := int32(len(f)) + 4
	if totalLen < 4 {
		return fmt.Errorf("totalLen of payload is out of range: %d", totalLen)
	}
	err := binary.Write(w, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}
	n, err := w.Write([]byte(f))
	if err != nil {
		return err
	}
	if n != len(f) {
		return ErrShortWrite
	}

	return nil
}

func (p *myFrameCodec) Decode(r io.Reader) (Payload, error) {
	var totalLen int32
	err := binary.Read(r, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}
	if totalLen < 4 {
		return nil, fmt.Errorf("totalLen is out of range: %x", totalLen)
	}

	buf := make([]byte, totalLen-4)
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	if n != int(totalLen-4) {
		return nil, ErrShortRead
	}

	return Payload(buf), nil
}
