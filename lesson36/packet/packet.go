package packet

import (
	"bytes"
	"fmt"
)

const (
	// CommandConn 0x01， 连接请求包
	CommandConn = iota + 0x01
	// CommandSubmit 0x02， 消息请求包
	CommandSubmit
)

const (
	// CommandConnAck   0x81，连接请求的响应包
	CommandConnAck = iota + 0x80
	// CommandSubmitAck 0x82，消息请求的响应包
	CommandSubmitAck
)

// Packet packet intf
type Packet interface {
	// Decode []byte -> struct
	Decode([]byte) error
	// Encode struct -> []byte
	Encode() ([]byte, error)
}

// Submit packet submit
type Submit struct {
	ID      string
	Payload []byte
}

// Decode submit decode
func (s *Submit) Decode(pktBody []byte) error {
	s.ID = string(pktBody[:8])
	s.Payload = pktBody[8:]
	return nil
}

// Encode submit encode
func (s *Submit) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(s.ID[:8]), s.Payload}, nil), nil
}

// SubmitAck packet submitAck
type SubmitAck struct {
	ID     string
	Result uint8
}

// Decode SubmitAck decode method
func (s *SubmitAck) Decode(pktBody []byte) error {
	s.ID = string(pktBody[:8])
	s.Result = uint8(pktBody[8])
	return nil
}

// Encode SubmitAck encode method
func (s *SubmitAck) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(s.ID[:8]), {s.Result}}, nil), nil
}

// Decode decode
func Decode(packet []byte) (Packet, error) {
	commandID := packet[0]
	pktBody := packet[1:]

	switch commandID {
	case CommandConn:
		return nil, nil
	case CommandConnAck:
		return nil, nil
	case CommandSubmit:
		s := Submit{}
		err := s.Decode(pktBody)
		if err != nil {
			return nil, err
		}
		return &s, nil
	case CommandSubmitAck:
		s := SubmitAck{}
		err := s.Decode(pktBody)
		if err != nil {
			return nil, err
		}
		return &s, nil
	default:
		return nil, fmt.Errorf("unknown commandID [%d]", commandID)
	}
}

// Encode encode
func Encode(p Packet) ([]byte, error) {
	var commandID uint8

	switch t := p.(type) {
	case *Submit:
		commandID = CommandSubmit
	case *SubmitAck:
		commandID = CommandSubmitAck
	default:
		return nil, fmt.Errorf("unknown type [%s]", t)
	}

	pktBody, err := p.Encode()

	if err != nil {
		return nil, err
	}

	return bytes.Join([][]byte{{commandID}, pktBody}, nil), nil
}
