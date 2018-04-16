package message

import (
	"encoding/binary"
	"io"
)

// ProofSize is the Cuckoo proof size.
const ProofSize = 42

// Proof is the proof of work.
type Proof struct {
	// Nonces are nonces.
	Nonces []uint32
}

//  Read reads the proof.
func (v *Proof) Read(r io.Reader) error {
	v.Nonces = make([]uint32, ProofSize)
	for i := 0; i < ProofSize; i++ {
		if err := binary.Read(r, binary.BigEndian, &v.Nonces[i]); err != nil {
			return err
		}
	}
	return nil
}
