package commands

import (
	"encoding/hex"
	"github.com/FigureTechnologies/extkey/pkg/encryption/extkey"
	"github.com/FigureTechnologies/extkey/pkg/encryption/pbkey"
	"github.com/FigureTechnologies/extkey/pkg/util"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cmdDecode = &cobra.Command{
	Use:   "decode [xprv..|xpub..]",
	Short: "Decode an xprv/xpub extended key",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hrp, err := cmd.PersistentFlags().GetString("hrp")
		if err != nil {
			return err
		}
		xkey := strings.TrimSpace(args[0])
		info, err := extkey.DecodeExtKey(hrp, xkey)
		if err != nil {
			return err
		}
		format := strings.TrimSpace(cmd.Flag("format").Value.String())
		return util.Display(info, format, os.Stdout)
	},
}

var cmdDecodeProto = &cobra.Command{
	Use:   "proto [--pub] [hex..]",
	Short: "Parse a p8e proto prv key (or pub if --pub is specified)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		protoBz, err := hex.DecodeString(strings.TrimSpace(args[0]))
		if err != nil {
			return err
		}
		pub, err := cmd.PersistentFlags().GetBool("pub")
		if err != nil {
			return err
		}
		info, err := pbkey.DecodeProto(protoBz, pub)
		if err != nil {
			return err
		}
		format := strings.TrimSpace(cmd.Flag("format").Value.String())
		return util.Display(info, format, os.Stdout)
	},
}

func init() {
	addFlags(cmdDecode, flagHRP, flagFormat)

	cmdDecodeProto.PersistentFlags().Bool("pub", false, "Decode public version of the proto bytes")
	cmdDecode.AddCommand(cmdDecodeProto)

	CmdRoot.AddCommand(cmdDecode)
}
