/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"github.com/hyperledger/fabric/gossip/common"
	"google.golang.org/grpc"
)

// MessageCryptoService is the contract between the gossip component and the
// peer's cryptographic layer and is used by the gossip component to verify,
// and authenticate remote peers and data they send, as well as to verify
// received blocks from the ordering service.
type MessageCryptoService interface {

	// GetPKIidOfCert returns the PKI-ID of a peer's identity
	// If any error occurs, the method return nil
	// This method does not validate peerIdentity.
	// This validation is supposed to be done appropriately during the execution flow.
	GetPKIidOfCert(peerIdentity PeerIdentityType) common.PKIidType

	// VerifyBlock returns nil if the block is properly signed, and the claimed seqNum is the
	// sequence number that the block's header contains.
	// else returns error
	VerifyBlock(chainID common.ChainID, seqNum uint64, signedBlock []byte) error

	// Sign signs msg with this peer's signing key and outputs
	// the signature if no error occurred.
	Sign(msg []byte) ([]byte, error)

	// Verify checks that signature is a valid signature of message under a peer's verification key.
	// If the verification succeeded, Verify returns nil meaning no error occurred.
	// If peerIdentity is nil, then the verification fails.
	Verify(peerIdentity PeerIdentityType, signature, message []byte) error

	// VerifyByChannel checks that signature is a valid signature of message
	// under a peer's verification key, but also in the context of a specific channel.
	// If the verification succeeded, Verify returns nil meaning no error occurred.
	// If peerIdentity is nil, then the verification fails.
	VerifyByChannel(chainID common.ChainID, peerIdentity PeerIdentityType, signature, message []byte) error

	// ValidateIdentity validates the identity of a remote peer.
	// If the identity is invalid, revoked, expired it returns an error.
	// Else, returns nil
	ValidateIdentity(peerIdentity PeerIdentityType) error
}

// PeerIdentityType is the peer's certificate
type PeerIdentityType []byte

// PeerSuspector returns whether a peer with a given identity is suspected
// as being revoked, or its CA is revoked
type PeerSuspector func(identity PeerIdentityType) bool

// PeerSecureDialOpts returns the gRPC DialOptions to use for connection level
// security when communicating with remote peer endpoints
type PeerSecureDialOpts func() []grpc.DialOption

// PeerSignature defines a signature of a peer
// on a given message
type PeerSignature struct {
	Signature    []byte
	Message      []byte
	PeerIdentity PeerIdentityType
}
