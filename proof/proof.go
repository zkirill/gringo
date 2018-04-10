// Package proof is concerned with proof of work.
package proof

// Proof is the proof of work.
type Proof struct {
	// Nonces are nonces.
	Nonces []uint32
	/// ProofSize is the size of the proof.
	ProofSize int
}
