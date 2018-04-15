package message

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/golang/glog"
)

// GetPeerAddrs is a request for peer addresses.
type GetPeerAddrs struct {
	// Capabilities filters peers by peer capabilities.
	Capabilities Capabilities
}

// Write writes request to get peer addresses.
func (v GetPeerAddrs) Write(w io.Writer) error {
	// Header.
	var h Header
	if err := h.Write(MsgTypeGetPeerAddrs, 4, w); err != nil {
		return fmt.Errorf("could not write header for ping message: %v", err)
	}
	// Body
	if err := binary.Write(w, binary.BigEndian, UnknownCapabilities); err != nil {
		return fmt.Errorf("could not write capabilities: %v", err)
	}
	return nil
}

// PeerAddrs contains peer addresses.
type PeerAddrs struct {
	// Peers represents peers.
	Peers []SockAddr
}

// Read reads peer addresses.
func (v *PeerAddrs) Read(r io.Reader) error {
	var len uint32
	// Length.
	if err := binary.Read(r, binary.BigEndian, &len); err != nil {
		return err
	}
	glog.Infof("reading %v peers", len)
	p := make([]SockAddr, len)
	for i := uint32(0); i < len; i++ {
		if err := p[i].Read(r); err != nil {
			return err
		}
	}
	v.Peers = p
	return nil
}

// Write writes peer addresses.
func (v PeerAddrs) Write(w io.Writer) error {
	// Length.
	if err := binary.Write(w, binary.BigEndian, uint32(len(v.Peers))); err != nil {
		return err
	}
	for _, p := range v.Peers {
		if err := p.Write(w); err != nil {
			return err
		}
	}
	return nil
}
