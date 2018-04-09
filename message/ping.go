package message

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Ping is the "ping" message.
type Ping struct {
	//  TotalDifficulty is the total difficulty accumulated by the user agent. It may be used to check whether sync may be needed.
	TotalDifficulty uint64
	// Height is the total height.
	Height uint64
}

// Read populates the ping with values from the reader.
func (p *Ping) Read(r io.Reader) error {
	// Total difficult.
	if err := binary.Read(r, binary.BigEndian, &p.TotalDifficulty); err != nil {
		return fmt.Errorf("could not read total difficulty: %v", err)
	}
	// Height.
	if err := binary.Read(r, binary.BigEndian, &p.Height); err != nil {
		return fmt.Errorf("could not read height: %v", err)
	}
	return nil
}

// Write writes the ping values to the writer.
// Set pong to true if message type should be set to "pong".
func (p *Ping) Write(pong bool, w io.Writer) error {
	// Send ping and pong are currently identical we can set the message type here.
	var msgType MsgType
	if pong {
		msgType = MsgTypePong
	} else {
		msgType = MsgTypePing
	}
	// Header.
	var h Header
	if err := h.Write(msgType, 16, w); err != nil {
		return fmt.Errorf("could not write header for ping message: %v", err)
	}
	// Total difficulty.
	if err := binary.Write(w, binary.BigEndian, &p.TotalDifficulty); err != nil {
		return fmt.Errorf("could not write total difficulty: %v", err)
	}
	// Height.
	if err := binary.Write(w, binary.BigEndian, &p.Height); err != nil {
		return fmt.Errorf("could not write height: %v", err)
	}
	return nil
}
