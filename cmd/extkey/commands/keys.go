package commands

import (
	"fmt"
	"math/big"

	bip32 "github.com/tyler-smith/go-bip32"
)

func DeriveChildKey(parentKey *bip32.Key, path string) (*bip32.Key, error) {
	childKey := parentKey
	bip44Indexes, bip44Harden, err := parseBIP44(path)
	if err != nil {
		panic(err)
	}
	for i, childIndex := range bip44Indexes {
		if bip44Harden[i] {
			childIndex |= bip32.FirstHardenedChild
		}
		childKey, err = childKey.NewChildKey(childIndex)
		if err != nil {
			return nil, err
		}
	}
	return childKey, nil
}

type SomeKey struct {
	Seed      string        `json:"seed,omitempty" yaml:"seed,omitempty"`
	Mnemonic  string        `json:"mnemonic,omitempty" yaml:"mnemonic,omitempty"`
	Hrp       string        `json:"hrp,omitempty" yaml:"hrp,omitempty"`
	RootKey   *ExtKeyData   `json:"rootKey,omitempty" yaml:"rootKey,omitempty"`
	ChildKeys []*ExtKeyData `json:"childKey,omitempty" yaml:"childKey,omitempty"`
}

type ExtKeyData struct {
	Address     string       `json:"address" yaml:"address"`
	Path        string       `json:"path,omitempty" yaml:"path,omitempty"`
	PrivateKey  InnerKeyData `json:"privateKey" yaml:"privateKey"`
	PublicKey   InnerKeyData `json:"publicKey" yaml:"publicKey"`
	Depth       byte         `json:"depth" yaml:"depth"`
	DepthLoc    string       `json:"depthLoc" yaml:"depthLoc"`
	Chaincode   string       `json:"chaincode" yaml:"chaincode"`
	Fingerprint string       `json:"fingerprint" yaml:"fingerprint"`
}

type InnerKeyData struct {
	Base58 string `json:"base58" yaml:"base58"`
	Bytes  string `json:"bytes" yaml:"bytes,flow"`
	BigInt string `json:"bigInt" yaml:"bigInt"`
}

func NewInnerKeyData(key *bip32.Key) InnerKeyData {
	bytes, _ := key.Serialize()
	return InnerKeyData{
		Base58: key.B58Serialize(),
		Bytes:  fmt.Sprintf("%X", bytes),
		BigInt: (&big.Int{}).SetBytes(key.Key).String(),
	}
}

func NewExtKeyData(key *bip32.Key, hrp, path string) *ExtKeyData {
	data := &ExtKeyData{
		Address:     toAddress(hrp, key),
		Path:        path,
		PublicKey:   NewInnerKeyData(key.PublicKey()),
		Depth:       key.Depth,
		DepthLoc:    depthString(key.Depth),
		Chaincode:   fmt.Sprintf("%X", key.ChainCode),
		Fingerprint: fmt.Sprintf("%X", key.FingerPrint),
	}
	if key.IsPrivate {
		data.PrivateKey = NewInnerKeyData(key)
	}
	return data
}
