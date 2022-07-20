package commands

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"io"
	"os"
	"strings"
)

var CmdGenerate = &cobra.Command{
	Use: "gen",
	RunE: func(cmd *cobra.Command, args []string) error {
		hrp := strings.TrimSpace(cmd.Flag("hrp").Value.String())
		hdPath := strings.TrimSpace(cmd.Flag("hd-path").Value.String())
		seed := strings.TrimSpace(cmd.Flag("seed").Value.String())
		formatter, err := formatize(strings.TrimSpace(cmd.Flag("format").Value.String()))
		if err != nil {
			return err
		}
		return generate(hrp, hdPath, seed, formatter, os.Stdout)
	},
}

func init() {
	addFlags(CmdGenerate, flagHRP, flagFormat, flagHDPath)
}

func generate(hrp, hdPath, seed string, formatter Formatter, w io.Writer) error {
	var seedBz []byte
	var err error
	if seed != "" {
		seedBz, err = base64.URLEncoding.DecodeString(seed)
		if err != nil {
			return err
		}
	}
	key, err := GenerateExtKey(hrp, hdPath, seedBz)
	if err != nil {
		return err
	}
	output, err := formatter(key)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%s\n", output)
	return nil
}

func GenerateExtKey(hrp, hdPath string, seedBz []byte) (someKey, error) {
	var seed []byte
	var err error
	var mnemonic string
	if seedBz == nil {
		entropy, err := bip39.NewEntropy(256)
		if err != nil {
			return someKey{}, err
		}
		mnemonic, err = bip39.NewMnemonic(entropy)
		if err != nil {
			return someKey{}, err
		}
		seed = bip39.NewSeed(mnemonic, "")
	} else {
		seed = seedBz
	}

	rootKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return someKey{}, err
	}
	key := someKey{
		Hrp:      hrp,
		Seed:     base64.URLEncoding.EncodeToString(seed),
		Mnemonic: mnemonic,
		RootKey:  NewExtKeyData(rootKey, hrp),
	}
	if hdPath != "" {
		childKey, err := DeriveChildKey(rootKey, hdPath)
		if err != nil {
			return someKey{}, err
		}
		key.ChildKey = NewExtKeyData(childKey, hrp)
	}
	return key, nil
}
