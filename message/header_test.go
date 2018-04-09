package message

import (
	"bytes"
	"testing"
)

func TestNewHeader(t *testing.T) {
	h, err := NewHeader(MsgTypeHand, 1)
	// Make a new header.
	if err != nil {
		t.Fatal(err)
	}
	// Check the length of the header.
	if h.Len() != HeaderLen {
		t.Errorf("wrong header length: expecting %v, got %v", HeaderLen, h.Len())
	}
}
func TestReadBadHeader(t *testing.T) {
	// Try reading headers that are too short.
	for i := 0; i < HeaderLen; i++ {
		b := make([]byte, i)
		r := bytes.NewReader(b)
		h := Header{}
		err := h.Read(r)
		if err == nil {
			t.Errorf("did not return error on bad header")
		}
	}
}

func TestWriteHeader(t *testing.T) {
	var b bytes.Buffer
	var h Header
	if err := h.Write(MsgTypeHand, 0, &b); err != nil {
		t.Error(err)
	}
}

func TestReadGoodHeader(t *testing.T) {
	msgType := MsgTypeHand
	var msgLen uint64 = 1
	b, _ := NewHeader(msgType, msgLen)
	h := Header{}
	// Read from bytes.
	if err := h.Read(bytes.NewReader(b.Bytes())); err != nil {
		t.Error(err)
	}
	// Magic 1.
	if h.Magic1 != Magic1 {
		t.Errorf("wrong first magic byte: expecting %v, got %v", Magic1, h.Magic1)
	}
	// Magic 2.
	if h.Magic2 != Magic2 {
		t.Errorf("wrong second magic byte: expecting %v, got %v", Magic2, h.Magic2)
	}
	// Message type.
	if h.MsgType != msgType {
		t.Errorf("wrong message type: expecting %v, got %v", msgType, h.MsgType)
	}
	// Message length.
	if h.Length != msgLen {
		t.Errorf("wrong message length: expecting %v, got %v", msgLen, h.Length)
	}
}
