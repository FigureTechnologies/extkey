package keys

import (
	"encoding/base64"
	"github.com/FigureTechnologies/extkey/pkg/encryption/extkey"
	"github.com/FigureTechnologies/extkey/pkg/encryption/types"
	"github.com/FigureTechnologies/extkey/pkg/util"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func NewSeedFromEnvOrPrompt() ([]byte, error) {
	var mnemonic string
	mnemonic, err := util.EnvOrSecret("mnemonic")
	if err != nil {
		return nil, err
	}

	var passphrase string
	passphrase, err = util.EnvOrSecret("passphrase")
	if err != nil {
		return nil, err
	}
	return bip39.NewSeed(mnemonic, passphrase), nil
}

func EncodeSomeKeysFromSeed(paths []string, hrp, seedB64 string) (*types.SomeKey, error) {
	var err error
	seedBz, err := base64.URLEncoding.DecodeString(seedB64)
	if err != nil {
		return nil, err
	}

	rootKey, err := bip32.NewMasterKey(seedBz)
	if err != nil {
		return nil, err
	}
	key := &types.SomeKey{
		Seed:     base64.URLEncoding.EncodeToString(seedBz),
		Mnemonic: "",
		Hrp:      hrp,
		RootKey:  extkey.NewExtKeyData(rootKey, hrp, ""),
	}
	for _, path := range paths {
		var childKey *bip32.Key
		childKey, err = extkey.DeriveChildKey(rootKey, path)
		if err != nil {
			return nil, err
		}
		key.ChildKeys = append(key.ChildKeys, extkey.NewExtKeyData(childKey, hrp, path))
	}
	return key, nil
}
