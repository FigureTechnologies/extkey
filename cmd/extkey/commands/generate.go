package commands

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	bip32 "github.com/tyler-smith/go-bip32"
	bip39 "github.com/tyler-smith/go-bip39"
)

var CmdGenerate = &cobra.Command{
	Use: "gen",
	RunE: func(cmd *cobra.Command, args []string) error {
		hrp := strings.TrimSpace(cmd.Flag("hrp").Value.String())
		if hrp == "" {
			return fmt.Errorf("--hrp is required")
		}
		hdPaths, err := cmd.PersistentFlags().GetStringArray("hd-path")
		if err != nil {
			return err
		}
		if len(hdPaths) == 0 {
			return fmt.Errorf("--hd-path is required")
		}
		var seed string
		if cmd.Flag("seed") != nil {
			seed = strings.TrimSpace(cmd.Flag("seed").Value.String())
		}

		formatter, err := formatize(strings.TrimSpace(cmd.Flag("format").Value.String()))
		if err != nil {
			return err
		}
		return generate(hrp, seed, hdPaths, formatter, os.Stdout)
	},
}

func init() {
	addFlags(CmdGenerate, flagHRP, flagFormat, flagHDPath, flagSeed)
}

func generate(hrp, seed string, paths []string, formatter Formatter, w io.Writer) error {
	var seedBz []byte
	var err error
	if seed != "" {
		seedBz, err = base64.URLEncoding.DecodeString(seed)
		if err != nil {
			return err
		}
	}
	key, err := GenerateExtKey(hrp, paths, seedBz)
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

func GenerateExtKey(hrp string, paths []string, seedBz []byte) (SomeKey, error) {
	var seed []byte
	var err error
	var mnemonic string
	if seedBz == nil {
		var entropy []byte
		entropy, err = bip39.NewEntropy(256)
		if err != nil {
			return SomeKey{}, err
		}
		mnemonic, err = bip39.NewMnemonic(entropy)
		if err != nil {
			return SomeKey{}, err
		}
		seed = bip39.NewSeed(mnemonic, "")
	} else {
		seed = seedBz
	}

	rootKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return SomeKey{}, err
	}
	key := SomeKey{
		Hrp:      hrp,
		Seed:     base64.URLEncoding.EncodeToString(seed),
		Mnemonic: mnemonic,
		RootKey:  NewExtKeyData(rootKey, hrp, ""),
	}
	for _, path := range paths {
		childKey, err := DeriveChildKey(rootKey, path)
		if err != nil {
			return SomeKey{}, err
		}
		key.ChildKeys = append(key.ChildKeys, NewExtKeyData(childKey, hrp, path))
	}
	return key, nil
}
