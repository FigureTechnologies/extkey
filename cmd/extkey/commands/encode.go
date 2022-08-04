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

var CmdEncode = &cobra.Command{
	Use: "encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, err := formatize(cmd.Flag("format").Value.String())
		if err != nil {
			return err
		}
		pHrp := strings.TrimSpace(cmd.Flag("hrp").Value.String())
		if pHrp == "" {
			return fmt.Errorf("hrp is required")
		}
		hdPath := strings.TrimSpace(cmd.Flag("hd-path").Value.String())
		seedB64 := strings.TrimSpace(cmd.Flag("seed").Value.String())
		return encode(hdPath, format, os.Stdout, pHrp, seedB64)
	},
}

func init() {
	addFlags(CmdEncode, flagFormat, flagHDPath, flagHRP, flagSeed)
}

func encode(path string, formatter Formatter, w io.Writer, hrp, seedB64 string) error {
	var seed []byte
	var err error
	if seedB64 == "" {
		var mnemonic string
		mnemonic, err = envOrSecret("mnemonic")
		if err != nil {
			return err
		}

		var passphrase string
		passphrase, err = envOrSecret("passphrase")
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
	key := SomeKey{
		Seed:     base64.URLEncoding.EncodeToString(seed),
		Mnemonic: "",
		Hrp:      hrp,
		RootKey:  NewExtKeyData(rootKey, hrp),
	}
	if hdPath != "" {
		var childKey *bip32.Key
		childKey, err = DeriveChildKey(rootKey, path)
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
