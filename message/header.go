package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// HeaderLen is the expected length of the header.
const HeaderLen int = 11

const (
	// Magic1 is the first magic byte.
	Magic1 uint8 = 0x1e
	// Magic2 is the second magic byte.
	Magic2 uint8 = 0xc5
)

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

// Write writes the header values to the writer.
func (h *Header) Write(msgType MsgType, msgLen uint64, w io.Writer) error {
	// Magic 1.
	if err := binary.Write(w, binary.BigEndian, Magic1); err != nil {
		return fmt.Errorf("could not write first magic byte: %v", err)
	}
	// Magic 2.
	if err := binary.Write(w, binary.BigEndian, Magic2); err != nil {
		return fmt.Errorf("could not write second magic byte: %v", err)
	}
	// Type of message.
	if err := binary.Write(w, binary.BigEndian, msgType); err != nil {
		return fmt.Errorf("could not write message type: %v", err)
	}
	// Length of message body.
	if err := binary.Write(w, binary.BigEndian, msgLen); err != nil {
		return fmt.Errorf("could not write message body length: %v", err)
	}
	return nil
}

// NewHeader returns new header.
func NewHeader(msgType MsgType, msgLen uint64) (bytes.Buffer, error) {
	var b bytes.Buffer
	// Magic 1.
	if err := binary.Write(&b, binary.BigEndian, Magic1); err != nil {
		return bytes.Buffer{}, err
	}
	// Magic 2.
	if err := binary.Write(&b, binary.BigEndian, Magic2); err != nil {
		return bytes.Buffer{}, err
	}
	// Type of message.
	if err := binary.Write(&b, binary.BigEndian, msgType); err != nil {
		return bytes.Buffer{}, err
	}
	// Length of message body.
	if err := binary.Write(&b, binary.BigEndian, msgLen); err != nil {
		return bytes.Buffer{}, err
	}
	return b, nil
}
