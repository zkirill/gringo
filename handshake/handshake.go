// Package handshake is responsible for the handshake.
package handshake

import (
	"github.com/golang/glog"
	"github.com/zkirill/mimblewimble-go/message"
)

// userAgent is the user agent of this client.
const userAgent = "mimblewimble-go 0.0.1"

// NewHandshake new handshake returns a new handshake.
func NewHandshake() ([]byte, error) {
	hand, err := message.NewHand()
	if err != nil {
		return nil, err
	}
	len := hand.Len()
	header, err := message.NewHeader(uint64(len))
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
