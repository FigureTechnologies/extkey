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

var CmdEncode = &cobra.Command{
	Use: "encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, err := formatize(cmd.Flag("format").Value.String())
		if err != nil {
			return err
		}
		hrp := strings.TrimSpace(cmd.Flag("hrp").Value.String())
		if hrp == "" {
			return fmt.Errorf("hrp is required")
		}
		hdPath := strings.TrimSpace(cmd.Flag("hd-path").Value.String())
		if hdPath == "" {
			return fmt.Errorf("hd-path is required")
		}
		return encode(hdPath, format, os.Stdout)
	},
}

func init() {
	addFlags(CmdEncode, flagFormat, flagHDPath, flagHRP)
}

func encode(path string, formatter Formatter, w io.Writer) error {
	mnemonic, err := envOrSecret("mnemonic")
	if err != nil {
		return err
	}

	passphrase, err := envOrSecret("passphrase")
	if err != nil {
		return err
	}

	seed := bip39.NewSeed(mnemonic, passphrase)
	rootKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return err
	}
	key := someKey{
		Seed:     seed,
		Mnemonic: mnemonic,
		Hrp:      hrp,
		RootKey:  NewExtKeyData(rootKey, hrp),
	}
	if hdPath != "" {
		childKey, err := DeriveChildKey(rootKey, path)
		if err != nil {
			return err
		}
		key.ChildKey = NewExtKeyData(childKey, hrp)
	}
	output, err := formatter(key)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%s\n", output)
	return nil
}
