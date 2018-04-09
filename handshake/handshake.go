// Package handshake is responsible for the handshake.
package handshake

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/golang/glog"
)

// userAgent is the user agent of this client.
const userAgent = "mimblewimble-go 0.0.1"

// HeaderLen is the expected length of the header.
const HeaderLen int = 11

// MsgType is the type of the message.
type MsgType uint8

const (
	// MsgTypeError represents an error.
	MsgTypeError MsgType = iota
	// MsgTypeHand represents the first part "hand" of the handshake.
	MsgTypeHand
	// MsgTypeShake represents the second part "shake" of the handshake.
	MsgTypeShake
)

// ProtocolVersion is the network protocol version.
type ProtocolVersion uint32

const (
	// ProtocolVersion1 is the current network protocol version.
	ProtocolVersion1 ProtocolVersion = 1
)

// Capabilities represents the capabilities of the client.
type Capabilities uint32

const (
	// UnknownCapabilities represents capabilities that are unknown.
	UnknownCapabilities Capabilities = 0 << 0
)

// Shake is the second part of the handshake.
type Shake struct {
	// Version is the version of the network on which is the sender.
	Version ProtocolVersion
	// Capabilities represents client capabilities of the sender.
	Capabilities Capabilities
	// Hash is the hash of the genesis.
	Hash [32]uint8
	// Total difficulty is the current total difficulty according to the sender.
	TotalDifficulty uint64
	// UserAgent is the user agent of the sender.
	UserAgent string
}

// Read populates the shake with values from the reader.
func (s *Shake) Read(r io.Reader) error {
	// Version.
	if err := binary.Read(r, binary.BigEndian, &s.Version); err != nil {
		return fmt.Errorf("could not read version: %v", err)
	}
	// Capabilities.
	if err := binary.Read(r, binary.BigEndian, &s.Capabilities); err != nil {
		return fmt.Errorf("could not read capabilities: %v", err)
	}
	// Total difficulty.
	if err := binary.Read(r, binary.BigEndian, &s.TotalDifficulty); err != nil {
		return fmt.Errorf("could not read total difficulty: %v", err)
	}
	// User agent length.
	var len uint64
	if err := binary.Read(r, binary.BigEndian, &len); err != nil {
		return fmt.Errorf("could not read user agent length: %v", err)
	}
	// User agent.
	agent := make([]byte, len)
	if err := binary.Read(r, binary.BigEndian, &agent); err != nil {
		return fmt.Errorf("could not read user agent: %v", err)
	}
	s.UserAgent = string(agent)
	// Genesis hash.
	if err := binary.Read(r, binary.BigEndian, &s.Hash); err != nil {
		return fmt.Errorf("could not read genesis hash: %v", err)
	}
	return nil
}

// Header is the message header.
type Header struct {
	// Magic1 is the first magic byte.
	Magic1 uint8
	// Magic2 is the second magic byte.
	Magic2 uint8
	// MsgType is the type of the message.
	MsgType MsgType
	// Length is the length of the message body (does not include the length of the header).
	Length uint64
}

// Read populates the header with the contents from the reader.
func (h *Header) Read(r io.Reader) error {
	// Magic 1.
	if err := binary.Read(r, binary.BigEndian, &h.Magic1); err != nil {
		return fmt.Errorf("could not read first magic byte: %v", err)
	}
	// Magic 2.
	if err := binary.Read(r, binary.BigEndian, &h.Magic2); err != nil {
		return fmt.Errorf("could not read second magic byte: %v", err)
	}
	// Message type.
	if err := binary.Read(r, binary.BigEndian, &h.MsgType); err != nil {
		return fmt.Errorf("could not read message type: %v", err)
	}
	// Message length.
	if err := binary.Read(r, binary.BigEndian, &h.Length); err != nil {
		return fmt.Errorf("could not read message body length: %v", err)
	}
	return nil
}

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
	if err := binary.Write(&b, binary.BigEndian, ProtocolVersion1); err != nil {
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
	ua := []byte(userAgent)
	if err := binary.Write(&b, binary.BigEndian, uint64(len(ua))); err != nil {
		return bytes.Buffer{}, err
	}
	// Now send the user agent.
	if err := binary.Write(&b, binary.BigEndian, ua); err != nil {
		return bytes.Buffer{}, err
	}
	// Genesis hash for test network 2.
	gen := [32]uint8{51, 70, 246, 60, 245, 178, 94, 20, 173, 221, 136, 85, 226, 117, 87, 132, 229, 94, 97, 44, 213, 133, 97, 200, 202, 24, 215, 207, 108, 168, 111, 75}
	if err := binary.Write(&b, binary.BigEndian, gen); err != nil {
		return bytes.Buffer{}, err
	}
	return b, nil
}
