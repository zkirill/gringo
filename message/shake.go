package message

import (
	"encoding/binary"
	"fmt"
	"io"
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
