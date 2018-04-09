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
