package commands

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/gogo/protobuf/proto"

	"github.com/FigureTechnologies/extkey/pkg/encryption/eckey"

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
	Address     string        `json:"address" yaml:"address"`
	Path        string        `json:"path,omitempty" yaml:"path,omitempty"`
	PrivateKey  *InnerKeyData `json:"privateKey" yaml:"privateKey"`
	PublicKey   *InnerKeyData `json:"publicKey" yaml:"publicKey"`
	Depth       byte          `json:"depth" yaml:"depth"`
	DepthLoc    string        `json:"depthLoc" yaml:"depthLoc"`
	Chaincode   string        `json:"chaincode" yaml:"chaincode"`
	Fingerprint string        `json:"fingerprint" yaml:"fingerprint"`
}

type InnerKeyData struct {
	Base58    string `json:"base58" yaml:"base58"`
	Base64    string `json:"base64" yaml:"base64"`
	Bytes     string `json:"bytes" yaml:"bytes,flow"`
	BigInt    string `json:"bigInt" yaml:"bigInt"`
	Proto     string `json:"proto" yaml:"proto"`
	ProtoJSON string `json:"protoJson" yaml:"protoJson"`
}

func NewInnerKeyDataFromBIP32(key *bip32.Key) (prvKey, pubKey *InnerKeyData) {
	if key.IsPrivate {
		prvKey = NewInnerKeyData(key.Key, key.B58Serialize(), false)
		pubKey = NewInnerKeyData(key.PublicKey().Key, key.PublicKey().B58Serialize(), true)
	} else {
		pubKey = NewInnerKeyData(key.Key, key.B58Serialize(), true)
	}
	return
}

func NewInnerKeyDataFromBTCECPub(key *btcec.PublicKey) (pubKey *InnerKeyData) {
	return NewInnerKeyData(key.SerializeCompressed(), "", true)
}

func NewInnerKeyDataFromBTCECPrv(key *btcec.PrivateKey) (prvKey, pubKey *InnerKeyData) {
	prvKey = NewInnerKeyData(key.Serialize(), "", false)
	pubKey = NewInnerKeyData(key.PubKey().SerializeCompressed(), "", true)
	return
}

func NewInnerKeyData(bz []byte, base58 string, compressed bool) *InnerKeyData {
	msg := &eckey.Key{
		KeyBytes:   bz,
		Type:       eckey.KeyType_ELLIPTIC,
		Curve:      eckey.KeyCurve_SECP256K1,
		Compressed: compressed,
	}
	protoBz, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	protoJS, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	b64 := base64.StdEncoding.EncodeToString(bz)
	return &InnerKeyData{
		Base58:    base58,
		Base64:    b64,
		Bytes:     fmt.Sprintf("%X", bz),
		BigInt:    (&big.Int{}).SetBytes(bz).String(),
		Proto:     fmt.Sprintf("%X", protoBz),
		ProtoJSON: string(protoJS),
	}
}

func NewExtKeyData(key *bip32.Key, hrp, path string) *ExtKeyData {
	prvKey, pubKey := NewInnerKeyDataFromBIP32(key)
	data := &ExtKeyData{
		Address:     toAddress(hrp, key),
		Path:        path,
		PublicKey:   pubKey,
		Depth:       key.Depth,
		DepthLoc:    depthString(key.Depth),
		Chaincode:   fmt.Sprintf("%X", key.ChainCode),
		Fingerprint: fmt.Sprintf("%X", key.FingerPrint),
	}
	if key.IsPrivate {
		data.PrivateKey = prvKey
	}
	return data
}
