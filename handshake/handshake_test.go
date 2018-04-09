package handshake

import (
	"testing"
)

func TestMakeHandshake(t *testing.T) {
	_, err := NewHandshake()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMakeHeader(t *testing.T) {
	h, err := NewHeader(1)
	// Make a new header.
	if err != nil {
		t.Fatal(err)
	}
	// Check the length of the header.
	if h.Len() != HeaderLen {
		t.Errorf("wrong header length: expecting %v, got %v", HeaderLen, h.Len())
	}
}
