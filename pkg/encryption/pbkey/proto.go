package pbkey

import (
	"fmt"
	"github.com/FigureTechnologies/extkey/pkg/encryption/eckey"
	"github.com/FigureTechnologies/extkey/pkg/encryption/extkey"
	"github.com/FigureTechnologies/extkey/pkg/keys"
	"github.com/btcsuite/btcd/btcec"
	"github.com/gogo/protobuf/proto"
)

func DecodeProto(bz []byte, isPub bool) (*extkey.DecodedProtoInfo, error) {
	msg := eckey.Key{}
	if err := proto.Unmarshal(bz, &msg); err != nil {
		return nil, err
	}

	var curve *btcec.KoblitzCurve
	switch msg.Curve {
	case eckey.KeyCurve_SECP256K1:
		curve = btcec.S256()
	default:
		panic("curve not found " + msg.Curve.String())
	}

	var pub *keys.InnerKeyData
	var prv *keys.InnerKeyData
	if isPub {
		pk, err := btcec.ParsePubKey(msg.KeyBytes, curve)
		if err != nil {
			return nil, err
		}
		if !curve.IsOnCurve(pk.X, pk.Y) {
			return nil, fmt.Errorf("invalid point for curve")
		}
		pub = extkey.NewInnerKeyDataFromBTCECPub(pk)
	} else {
		prvK, pubK := btcec.PrivKeyFromBytes(curve, msg.KeyBytes)
		if prvK == nil || pubK == nil {
			return nil, fmt.Errorf("invalid private key")
		}
		if !curve.IsOnCurve(prvK.X, prvK.Y) {
			return nil, fmt.Errorf("invalid point for private key")
		}
		if !curve.IsOnCurve(pubK.X, pubK.Y) {
			return nil, fmt.Errorf("invalid point for public key")
		}
		prv, pub = extkey.NewInnerKeyDataFromBTCECPrv(prvK)
	}
	info := &extkey.DecodedProtoInfo{
		PublicKey:  pub,
		PrivateKey: prv,
	}
	return info, nil
}
