package commands

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	bip32 "github.com/tyler-smith/go-bip32"
)

var CmdDecode = &cobra.Command{
	Use:   "decode [xprv..|xpub..]",
	Short: "Decode an xprv/xpub extended key",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hrp, err := cmd.PersistentFlags().GetString("hrp")
		if err != nil {
			return err
		}
		formatter, err := formatize(strings.TrimSpace(cmd.Flag("format").Value.String()))
		if err != nil {
			return err
		}
		extkey := strings.TrimSpace(args[0])
		return decode(hrp, extkey, os.Stdout, formatter)
	},
}

func init() {
	addFlags(CmdDecode, flagHRP, flagFormat)
}

func decode(hrp, xkey string, w io.Writer, formatter Formatter) error {
	key, err := bip32.B58Deserialize(xkey)
	if err != nil {
		return err
	}

	info := decodedKeyInfo{}
	info.Address = toAddress(hrp, key)
	info.XKey = SomeKey{
		Seed:      "",
		Mnemonic:  "",
		Hrp:       hrp,
		RootKey:   NewExtKeyData(key, hrp, ""),
		ChildKeys: nil,
	}

	output, err := formatter(info)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%s\n", output)
	return nil
}

type decodedKeyInfo struct {
	Address string  `json:"address" yaml:"address"`
	XKey    SomeKey `json:"xkey" yaml:"xkey"`
}
