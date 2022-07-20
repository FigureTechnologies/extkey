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
		seedB64 := strings.TrimSpace(cmd.Flag("seed").Value.String())
		return encode(hdPath, format, os.Stdout, seedB64)
	},
}

func init() {
	addFlags(CmdEncode, flagFormat, flagHDPath, flagHRP, flagSeed)
}

func encode(path string, formatter Formatter, w io.Writer, seedB64 string) error {
	var seed []byte
	var err error
	if seedB64 == "" {
		mnemonic, err := envOrSecret("mnemonic")
		if err != nil {
			return err
		}

		passphrase, err := envOrSecret("passphrase")
		if err != nil {
			return err
		}

		seed = bip39.NewSeed(mnemonic, passphrase)
	} else {
		seed, err = base64.URLEncoding.DecodeString(seedB64)
		if err != nil {
			return err
		}
	}

	rootKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return err
	}
	key := someKey{
		Seed:     base64.URLEncoding.EncodeToString(seed),
		Mnemonic: "",
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
