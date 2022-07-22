package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip32"
	"io"
	"os"
	"strings"
)

var CmdDecode = &cobra.Command{
	Use:  "decode",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		formatter, err := formatize(strings.TrimSpace(cmd.Flag("format").Value.String()))
		if err != nil {
			return err
		}
		extkey := strings.TrimSpace(args[0])
		decode(extkey, os.Stdout, formatter)
		return nil
	},
}

func init() {
	addFlags(CmdDecode, flagHRP, flagFormat)
}

func decode(xkey string, w io.Writer, formatter Formatter) error {
	key, err := bip32.B58Deserialize(xkey)
	if err != nil {
		return err
	}

	info := decodedKeyInfo{}
	info.Address = toAddress(hrp, key)
	info.XKey.Depth = key.Depth
	info.XKey.DepthLoc = depthString(key.Depth)
	info.XKey.Chaincode = fmt.Sprintf("%X", key.ChainCode)
	info.XKey.Fingerprint = fmt.Sprintf("%X", key.FingerPrint)

	if key.IsPrivate {
		info.XKey.Private = key.B58Serialize()
		info.XKey.Public = key.PublicKey().B58Serialize()
	} else {
		info.XKey.Public = key.B58Serialize()
	}

	output, err := formatter(info)
	fmt.Fprintf(w, "%s\n", output)
	return nil
}

type decodedKeyInfo struct {
	Address string `json:"address" yaml:"address"`
	XKey    struct {
		Private     string `json:"private,omitempty" yaml:"private,omitempty"`
		Public      string `json:"public" yaml:"public"`
		Depth       byte   `json:"depth" yaml:"depth"`
		DepthLoc    string `json:"depthLoc" yaml:"depthLoc"`
		Chaincode   string `json:"chaincode" yaml:"chaincode"`
		Fingerprint string `json:"fingerprint" yaml:"fingerprint"`
	} `json:"xkey" yaml:"xkey"`
}
