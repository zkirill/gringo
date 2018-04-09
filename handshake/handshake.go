// Package handshake is responsible for the handshake.
package handshake

import (
	"github.com/golang/glog"
	"github.com/zkirill/gringo/message"
)

// NewHandshake new handshake returns a new handshake.
func NewHandshake() ([]byte, error) {
	hand, err := message.NewHand()
	if err != nil {
		return nil, err
	}
	len := hand.Len()
	header, err := message.NewHeader(message.MsgTypeHand, uint64(len))
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
