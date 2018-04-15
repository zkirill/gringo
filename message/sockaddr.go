package message

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// SockAddr represents an address.
type SockAddr struct {
	// Addr represents the address.
	Addr net.IPAddr
	// Port is the port.
	Port uint16
}

// Read reads the sock address in a format that is sent by Grin.
func (v *SockAddr) Read(w io.Reader) error {
	var t uint8
	// Leading IP type.
	if err := binary.Read(w, binary.BigEndian, &t); err != nil {
		return err
	}
	if t != 0 {
		// TODO: Handle IPv6.
		return nil
	}
	ip := make([]uint8, 4)
	// IP v4 in the next 4 parts.
	if err := binary.Read(w, binary.BigEndian, &ip[0]); err != nil {
		return err
	}
	if err := binary.Read(w, binary.BigEndian, &ip[1]); err != nil {
		return err
	}
	if err := binary.Read(w, binary.BigEndian, &ip[2]); err != nil {
		return err
	}
	if err := binary.Read(w, binary.BigEndian, &ip[3]); err != nil {
		return err
	}
	v.Addr.IP = net.IPv4(ip[0], ip[1], ip[2], ip[3])
	// Port.
	if err := binary.Read(w, binary.BigEndian, &v.Port); err != nil {
		return fmt.Errorf("could not read port: %v", err)
	}
	return nil
}

// Write writes the sock address in a format that is understood by Grin.
func (v SockAddr) Write(w io.Writer) error {
	// Convert to v4.
	ip := v.Addr.IP.To4()
	if ip == nil {
		return fmt.Errorf("invalid IP address")
	}
	// Leading zero because this is IP v4.
	if err := binary.Write(w, binary.BigEndian, uint8(0)); err != nil {
		return err
	}
	// IP v4 in the next 4 parts.
	if err := binary.Write(w, binary.BigEndian, uint8(ip[0])); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, uint8(ip[1])); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, uint8(ip[2])); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, uint8(ip[3])); err != nil {
		return err
	}
	// Port.
	if err := binary.Write(w, binary.BigEndian, v.Port); err != nil {
		return fmt.Errorf("could not write port: %v", err)
	}
	return nil
}
