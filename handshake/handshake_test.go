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
	_, err := NewHeader(1)
	if err != nil {
		t.Fatal(err)
	}
}
