// Package handshake is responsible for the handshake.
package handshake

import (
	"bytes"
	"encoding/binary"

	"github.com/golang/glog"
)

const (
	// MsgTypeError represents an error.
	MsgTypeError = iota
	// MsgTypeHand represents the first part "hand" of the handshake.
	MsgTypeHand
	// MsgTypeShake represents the second part "shake" of the handshake.
	MsgTypeShake
)

const (
	// ProtocolVersion is the network protocol version.
	ProtocolVersion = 1
)

const (
	// UnknownCapabilities represents capabilities that are unknown.
	UnknownCapabilities uint8 = 0 << 0
)

// NewHeader returns new header.
func NewHeader(msgLen uint64) (bytes.Buffer, error) {
	var b bytes.Buffer
	// Magic 1.
	if err := binary.Write(&b, binary.BigEndian, uint8(0x1e)); err != nil {
		return bytes.Buffer{}, err
	}
	// Magic 2.
	if err := binary.Write(&b, binary.BigEndian, uint8(0xc5)); err != nil {
		return bytes.Buffer{}, err
	}
	// Type of message.
	if err := binary.Write(&b, binary.BigEndian, uint8(MsgTypeHand)); err != nil {
		return bytes.Buffer{}, err
	}
	// Length of message body.
	if err := binary.Write(&b, binary.BigEndian, msgLen); err != nil {
		return bytes.Buffer{}, err
	}
	return b, nil
}

// NewHandshake new handshake returns a new handshake.
func NewHandshake() ([]byte, error) {
	hand, err := NewHand()
	if err != nil {
		return nil, err
	}
	len := hand.Len()
	header, err := NewHeader(uint64(len))
	if err != nil {
		return nil, err
	}
	n, err := header.Write(hand.Bytes())
	if err != nil {
		return nil, err
	}
	glog.Infof("add %v bytes to buffer", n)
	return header.Bytes(), nil
}

// NewHand returns a new handshake.
func NewHand() (bytes.Buffer, error) {
	var b bytes.Buffer
	// Protocol version (u32).
	if err := binary.Write(&b, binary.BigEndian, uint32(ProtocolVersion)); err != nil {
		return bytes.Buffer{}, err
	}
	// Capabilities (u32).
	if err := binary.Write(&b, binary.BigEndian, uint32(0)); err != nil {
		return bytes.Buffer{}, err
	}
	// Nonce (u64).
	if err := binary.Write(&b, binary.BigEndian, uint64(1)); err != nil {
		return bytes.Buffer{}, err
	}
	// Total difficulty (u64).
	if err := binary.Write(&b, binary.BigEndian, uint64(1)); err != nil {
		return bytes.Buffer{}, err
	}
	// Sender address.
	// Leading zero because this is IP v4.
	if err := binary.Write(&b, binary.BigEndian, uint8(0)); err != nil {
		return bytes.Buffer{}, err
	}
	// Sender IP v4 in the next 4 parts.
	if err := binary.Write(&b, binary.BigEndian, uint8(0)); err != nil {
		return bytes.Buffer{}, err

	}
	if err := binary.Write(&b, binary.BigEndian, uint8(0)); err != nil {
		return bytes.Buffer{}, err
	}
	if err := binary.Write(&b, binary.BigEndian, uint8(0)); err != nil {
		return bytes.Buffer{}, err
	}
	if err := binary.Write(&b, binary.BigEndian, uint8(0)); err != nil {
		return bytes.Buffer{}, err
	}
	// Sender port.
	if err := binary.Write(&b, binary.BigEndian, uint16(13414)); err != nil {
		return bytes.Buffer{}, err
	}
	// Receiver address.
	// Leading zero because this is IP v4.
	if err := binary.Write(&b, binary.BigEndian, uint8(0)); err != nil {
		return bytes.Buffer{}, err
	}
	// IP v4 in the next 4 parts.
	if err := binary.Write(&b, binary.BigEndian, uint8(0)); err != nil {
		return bytes.Buffer{}, err
	}
	if err := binary.Write(&b, binary.BigEndian, uint8(0)); err != nil {
		return bytes.Buffer{}, err
	}
	if err := binary.Write(&b, binary.BigEndian, uint8(0)); err != nil {
		return bytes.Buffer{}, err
	}
	if err := binary.Write(&b, binary.BigEndian, uint8(0)); err != nil {
		return bytes.Buffer{}, err
	}
	// Receiver port.
	if err := binary.Write(&b, binary.BigEndian, uint16(13414)); err != nil {
		return bytes.Buffer{}, err
	}
	// User agent.
	// First we need to send the length.
	// TODO: Do not hardcode the length.
	if err := binary.Write(&b, binary.BigEndian, uint64(2)); err != nil {
		return bytes.Buffer{}, err
	}
	// Now send the user agent.
	if err := binary.Write(&b, binary.BigEndian, []byte("Go")); err != nil {
		return bytes.Buffer{}, err
	}
	// Genesis hash for test network 2.
	gen := [32]uint8{51, 70, 246, 60, 245, 178, 94, 20, 173, 221, 136, 85, 226, 117, 87, 132, 229, 94, 97, 44, 213, 133, 97, 200, 202, 24, 215, 207, 108, 168, 111, 75}
	if err := binary.Write(&b, binary.BigEndian, gen); err != nil {
		return bytes.Buffer{}, err
	}
	return b, nil
}
