package message

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/golang/glog"
)

type Locator struct {
	Hashes []Hash
}

func (v Locator) Write(w io.Writer) error {
	glog.Info("writing length")
	len := len(v.Hashes)
	if len == 0 {
		// Write genesis hash.
		if err := binary.Write(w, binary.BigEndian, uint8(1)); err != nil {
			return fmt.Errorf("could not write length: %v", err)
		}
		glog.Info("writing genesis hash")
		if err := binary.Write(w, binary.BigEndian, GenesisHash()); err != nil {
			return fmt.Errorf("could not write genesis hash: %v", err)
		}
		return nil
	}
	for i := 0; i < len; i++ {
		// Write more than one hash.
		if err := binary.Write(w, binary.BigEndian, uint8(len)); err != nil {
			return fmt.Errorf("could not write length: %v", err)
		}
		if err := binary.Write(w, binary.BigEndian, v.Hashes[i]); err != nil {
			return fmt.Errorf("could not write hash: %v", err)
		}
	}
	return nil
}
