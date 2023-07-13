// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package sr25519

import (
	"crypto/rand"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"

	"github.com/centrifuge/chainbridge-utils/crypto"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ crypto.Keypair = &Keypair{}

type Keypair struct {
	keyringPair *signature.KeyringPair
}

func GenerateKeypair(network uint16) (*Keypair, error) {
	data := make([]byte, 32)
	_, err := rand.Read(data)
	if err != nil {
		return nil, err
	}
	return NewKeypairFromSeed("//"+hexutil.Encode(data), network)
}

func NewKeypairFromSeed(seed string, network uint16) (*Keypair, error) {
	kp, err := signature.KeyringPairFromSecret(seed, network)
	return &Keypair{&kp}, err
}

func NewKeypairFromKRP(pair signature.KeyringPair) *Keypair {
	return &Keypair{&pair}
}

// AsKeyringPair returns the underlying KeyringPair
func (kp *Keypair) AsKeyringPair() *signature.KeyringPair {
	return kp.keyringPair
}

// Encode uses scale to encode underlying KeyringPair
func (kp *Keypair) Encode() []byte {
	out, _ := codec.Encode(kp.keyringPair)
	return out
}

// Decode initializes keypair by decoding input as a KeyringPair
func (kp *Keypair) Decode(in []byte) error {
	kp.keyringPair = &signature.KeyringPair{}
	return codec.Decode(in, kp.keyringPair)
}

// Address returns the ss58 formated address
func (kp *Keypair) Address() string {
	return kp.keyringPair.Address
}

// PublicKey returns the publickey encoded as a string
func (kp *Keypair) PublicKey() string {
	return hexutil.Encode(kp.keyringPair.PublicKey)
}
