package extkey

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/FigureTechnologies/extkey/pkg/encryption/eckey"
	"github.com/FigureTechnologies/extkey/pkg/encryption/types"
	"github.com/FigureTechnologies/extkey/pkg/util"
	"github.com/btcsuite/btcd/btcec"
	"github.com/gogo/protobuf/proto"
	"github.com/tyler-smith/go-bip32"
	"math/big"
)

func DeriveChildKey(parentKey *bip32.Key, path string) (*bip32.Key, error) {
	childKey := parentKey
	bip44Indexes, bip44Harden, err := util.ParseBIP44(path)
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

func NewInnerKeyDataFromBIP32(key *bip32.Key) (prvKey, pubKey *types.InnerKeyData) {
	if key.IsPrivate {
		prvKey = NewInnerKeyData(key.Key, key.B58Serialize(), false)
		pubKey = NewInnerKeyData(key.PublicKey().Key, key.PublicKey().B58Serialize(), true)
	} else {
		pubKey = NewInnerKeyData(key.Key, key.B58Serialize(), true)
	}
	return
}

func NewInnerKeyDataFromBTCECPub(key *btcec.PublicKey) (pubKey *types.InnerKeyData) {
	return NewInnerKeyData(key.SerializeCompressed(), "", true)
}

func NewInnerKeyDataFromBTCECPrv(key *btcec.PrivateKey) (prvKey, pubKey *types.InnerKeyData) {
	prvKey = NewInnerKeyData(key.Serialize(), "", false)
	pubKey = NewInnerKeyData(key.PubKey().SerializeCompressed(), "", true)
	return
}

func NewInnerKeyData(bz []byte, base58 string, compressed bool) *types.InnerKeyData {
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
	return &types.InnerKeyData{
		Base58:    base58,
		Base64:    b64,
		Bytes:     fmt.Sprintf("%X", bz),
		BigInt:    (&big.Int{}).SetBytes(bz).String(),
		Proto:     fmt.Sprintf("%X", protoBz),
		ProtoJSON: string(protoJS),
	}
}

func NewExtKeyData(key *bip32.Key, hrp, path string) *types.ExtKeyData {
	prvKey, pubKey := NewInnerKeyDataFromBIP32(key)
	data := &types.ExtKeyData{
		Address:     util.KeyToAddress(hrp, key),
		Path:        path,
		PublicKey:   pubKey,
		Depth:       key.Depth,
		DepthLoc:    util.DepthString(key.Depth),
		Chaincode:   fmt.Sprintf("%X", key.ChainCode),
		Fingerprint: fmt.Sprintf("%X", key.FingerPrint),
	}
	if key.IsPrivate {
		data.PrivateKey = prvKey
	}
	return data
}
