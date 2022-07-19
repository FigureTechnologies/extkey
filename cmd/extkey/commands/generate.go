package commands

import (
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
		formatter, err := formatize(strings.TrimSpace(cmd.Flag("format").Value.String()))
		if err != nil {
			return err
		}
		return generate(hrp, hdPath, formatter, os.Stdout)
	},
}

func init() {
	addFlags(CmdGenerate, flagHRP, flagFormat, flagHDPath)
}

func generate(hrp, hdPath string, formatter Formatter, w io.Writer) error {
	key, err := GenerateExtKey(hrp, hdPath)
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

func GenerateExtKey(hrp, hdPath string) (someKey, error) {
	seed, err := bip39.NewEntropy(256)
	if err != nil {
		return someKey{}, err
	}

	mnemonic, err := bip39.NewMnemonic(seed)
	if err != nil {
		return someKey{}, err
	}

	rootKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return someKey{}, err
	}
	key := someKey{
		Hrp:      hrp,
		Seed:     seed,
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
