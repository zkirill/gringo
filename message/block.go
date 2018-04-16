package message

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"github.com/golang/glog"
)

// GetHeaders requests block headers.
type GetHeaders struct {
	Locator Locator
}

// Write writes message for getting headers.
func (v GetHeaders) Write(w io.Writer) error {
	// Header.
	var h Header
	if err := h.Write(MsgTypeGetHeaders, 33, w); err != nil {
		return fmt.Errorf("could not write header for GetHeaders message: %v", err)
	}
	// Locator.
	if err := v.Locator.Write(w); err != nil {
		return fmt.Errorf("could not write locator: %v", err)
	}
	return nil
}

// BlockHeaders is a wrapper for headers.
type BlockHeaders struct {
	// Headers are headers.
	Headers []BlockHeader
}

// BlockHeaders reads headers.
func (v *BlockHeaders) Read(r io.Reader) error {
	var len uint16
	if err := binary.Read(r, binary.BigEndian, &len); err != nil {
		return err
	}
	glog.Infof("received %v block headers", len)
	v.Headers = make([]BlockHeader, len)
	for i := uint16(0); i < len; i++ {
		v.Headers[i].Read(r)
	}
	return nil
}

// BlockHeader is the MimbleWimble block header.
type BlockHeader struct {
	/// Version is the version of the block.
	Version uint16
	/// Height is the height of this block since the genesis block (height 0).
	Height uint64
	/// Previous is the hash of the block previous to this in the chain.
	Previous Hash
	/// Timestamp is the timestamp at at which the block was built.
	Timestamp time.Time
	/// TotalDifficulty is the total accumulated difficulty since the genesis block.
	TotalDifficulty uint64
	/// OutputRoot is the merklish root of all the commitments in the TxHashSet.
	OutputRoot Hash
	/// RangeProofRoot is the merklish root of all range proofs in the TxHashSet.
	RangeProofRoot Hash
	/// KernelRoot is the root of all transaction kernels in the TxHashSet.
	KernelRoot Hash
	/// TotalKernelOffset is the total accumulated sum of kernel offsets since genesis block.
	/// We can derive the kernel offset sum for *this* block from
	/// the total kernel offset of the previous block header.
	// https://github.com/mimblewimble/rust-secp256k1-zkp/blob/f41f661928ed177fc20e9cb23f23d940cef742bd/src/constants.rs#L23
	TotalKernelOffset [32]uint8
	/// Nonce is the nonce increment used to mine this block.
	Nonce uint64
	/// Proof of work data.
	ProofOfWork Proof
}

func (v *BlockHeader) Read(r io.Reader) error {
	// Version.
	if err := binary.Read(r, binary.BigEndian, &v.Version); err != nil {
		return err
	}
	// Height.
	if err := binary.Read(r, binary.BigEndian, &v.Height); err != nil {
		return err
	}
	// Hash of the previous block to this block in the chain.
	if err := binary.Read(r, binary.BigEndian, &v.Previous); err != nil {
		return err
	}
	// Timestamp.
	var ts int64
	if err := binary.Read(r, binary.BigEndian, &ts); err != nil {
		return err
	}
	v.Timestamp = time.Unix(ts, 0)
	// Total difficulty.
	if err := binary.Read(r, binary.BigEndian, &v.TotalDifficulty); err != nil {
		return err
	}
	// Output root.
	if err := binary.Read(r, binary.BigEndian, &v.OutputRoot); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &v.RangeProofRoot); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &v.KernelRoot); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &v.TotalKernelOffset); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &v.Nonce); err != nil {
		return err
	}
	// Proof of work.
	if err := v.ProofOfWork.Read(r); err != nil {
		return fmt.Errorf("could not read pow: %v", err)
	}
	return nil
}

type Block struct {
	// Header contains metadata and commitments to the rest of the data.
	Header BlockHeader
	// Inputs is the list of transaction inputs.
	Inputs []Input
	// Outputs is the list of transaction outputs.
	Outputs []Output
	// Kernels is a list of kernels.
	Kernels []TxKernel
}

func (v *Block) Read(r io.Reader) error {
	if err := v.Header.Read(r); err != nil {
		return fmt.Errorf("could not read block header: %v", err)
	}
	var inputsLen, outputsLen, kernelsLen uint64
	if err := binary.Read(r, binary.BigEndian, &inputsLen); err != nil {
		return fmt.Errorf("could not read inputs length: %v", err)
	}
	if err := binary.Read(r, binary.BigEndian, &outputsLen); err != nil {
		return fmt.Errorf("could not read outputs length: %v", err)
	}
	if err := binary.Read(r, binary.BigEndian, &kernelsLen); err != nil {
		return fmt.Errorf("could not read kernels length: %v", err)
	}
	return nil
}

type Input struct {
	Features OutputFeatures
	Commit   [33]uint8
}

type Output struct {
	Features OutputFeatures
	Commit   [33]uint8
	Proof    PangeProof
}

type OutputFeatures uint8

const (
	DefaultOutputFeatures  OutputFeatures = 0 << 0
	CoinbaseOutputFeatures OutputFeatures = 0 << 1
)

type PangeProof struct {
	Proof      [5134]uint8
	ProofLenth uint
}

type TxKernel struct {
}

// GetBlock requests block by hash.
func GetBlock(hash Hash, w io.Writer) error {
	// Header.
	var h Header
	if err := h.Write(MsgTypeGetBlock, 32, w); err != nil {
		return fmt.Errorf("could not write header for GetBlock message: %v", err)
	}
	// Block hash.
	if err := binary.Write(w, binary.BigEndian, &hash); err != nil {
		return fmt.Errorf("could not write hash: %v", err)
	}
	return nil
}
