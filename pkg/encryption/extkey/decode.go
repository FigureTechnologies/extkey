package extkey

import (
	"github.com/FigureTechnologies/extkey/pkg/keys"
	"github.com/FigureTechnologies/extkey/pkg/util"
	"github.com/tyler-smith/go-bip32"
)

type DecodedProtoInfo struct {
	PrivateKey *keys.InnerKeyData `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`
	PublicKey  *keys.InnerKeyData `json:"publicKey" yaml:"publicKey"`
}

type DecodedExtKeyInfo struct {
	Address string       `json:"address" yaml:"address"`
	XKey    keys.SomeKey `json:"xkey" yaml:"xkey"`
}

func DecodeExtKey(hrp, extKey string) (DecodedExtKeyInfo, error) {
	key, err := bip32.B58Deserialize(extKey)
	if err != nil {
		return DecodedExtKeyInfo{}, err
	}
	info := DecodedExtKeyInfo{}
	info.Address = util.KeyToAddress(hrp, key)
	info.XKey = keys.SomeKey{
		Seed:      "",
		Mnemonic:  "",
		Hrp:       hrp,
		RootKey:   NewExtKeyData(key, hrp, ""),
		ChildKeys: nil,
	}
	return info, nil
}
