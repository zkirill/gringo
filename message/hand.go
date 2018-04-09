package message

import (
	"bytes"
	"encoding/binary"
)

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
