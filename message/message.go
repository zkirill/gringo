// Package message handles messages that the user agent can send and receive.
package message

// MsgType is the type of the message.
type MsgType uint8

const (
	// MsgTypeError represents an error.
	MsgTypeError MsgType = iota
	// MsgTypeHand represents the first part "hand" of the handshake.
	MsgTypeHand
	// MsgTypeShake represents the second part "shake" of the handshake.
	MsgTypeShake
	// MsgTypePing is the "ping".
	MsgTypePing
	// MsgTypePong is the "pong".
	MsgTypePong
	// MsgTypeGetPeerAddrs represents a request for peers.
	MsgTypeGetPeerAddrs
	// MsgTypePeerAddrs represents a response containing peer addresses.
	MsgTypePeerAddrs
	// MsgTypeGetHeaders is a request for block headers.
	MsgTypeGetHeaders
	// MsgTypeHeader is a block header.
	MsgTypeHeader
	// MsgTypeHeaders are headers.
	MsgTypeHeaders
	// MsgTypeGetBlock is a request for block.
	MsgTypeGetBlock
	// MsgTypeBlock is a block.
	MsgTypeBlock
	// MsgTypeGetCompactBlock is a request for a compact block.
	MsgTypeGetCompactBlock
	// MsgTypeCompactBlock is a compact block.
	MsgTypeCompactBlock
	// MsgTypeStemTransaction is a stem transaction.
	// https://github.com/mimblewimble/grin/blob/e31205471404b3ffa9a0fc6ff3b0833c086c4b7a/doc/dandelion/dandelion.md
	MsgTypeStemTransaction
	// MsgTypeTransaction is a transaction.
	MsgTypeTransaction
	// MsgTypeTxHashSetRequest is a hash set request.
	MsgTypeTxHashSetRequest
	// MsgTypeTxHashSetArchive is a hash set archive.
	MsgTypeTxHashSetArchive
)

// userAgent is the user agent of this client.
const userAgent = "gringo 0.0.1"

// ProtocolVersion is the network protocol version.
type ProtocolVersion uint32

const (
	// ProtocolVersion1 is the current network protocol version.
	ProtocolVersion1 ProtocolVersion = 1
)

// Capabilities represents the capabilities of the client.
type Capabilities uint32

const (
	// UnknownCapabilities represents capabilities that are unknown.
	UnknownCapabilities Capabilities = 0 << 0
)
