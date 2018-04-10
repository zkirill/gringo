// Package block contains block and blockheader functionality.
package block

import (
	"time"

	"github.com/zkirill/gringo/proof"
)

// Header is the MimbleWimble block header.
type Header struct {
	/// Version is the version of the block.
	Version uint16
	/// Height is the height of this block since the genesis block (height 0).
	Height uint64
	/// Previous is the hash of the block previous to this in the chain.
	Previous [32]uint8
	/// Timestamp is the timestamp at at which the block was built.
	Timestamp time.Time
	/// TotalDifficulty is the total accumulated difficulty since the genesis block.
	TotalDifficulty uint32
	/// OutputRoot is the merklish root of all the commitments in the TxHashSet.
	OutputRoot [32]uint8
	/// RangeProofRoot is the merklish root of all range proofs in the TxHashSet.
	RangeProofRoot [32]uint8
	/// KernelRoot is the root of all transaction kernels in the TxHashSet.
	KernelRoot [32]uint8
	/// TotalKernelOffset is the total accumulated sum of kernel offsets since genesis block.
	/// We can derive the kernel offset sum for *this* block from
	/// the total kernel offset of the previous block header.
	// https://github.com/mimblewimble/rust-secp256k1-zkp/blob/f41f661928ed177fc20e9cb23f23d940cef742bd/src/constants.rs#L23
	TotalKernelOffset [32]uint8
	/// Nonce is the nonce increment used to mine this block.
	Nonce uint64
	/// Proof of work data.
	ProofOfWork proof.Proof
}
