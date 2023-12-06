package commands

import (
	"fmt"
	"github.com/FigureTechnologies/extkey/pkg/keys"
	"github.com/FigureTechnologies/extkey/pkg/util"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var CmdEncode = &cobra.Command{
	Use:   "encode",
	Short: "Output an xprv/xpub from a mnemonic, passphrase, hd-path, and hrp",
	RunE: func(cmd *cobra.Command, args []string) error {
		pHrp := strings.TrimSpace(cmd.Flag("hrp").Value.String())
		if pHrp == "" {
			return fmt.Errorf("hrp is required")
		}
		hdPaths, err := cmd.PersistentFlags().GetStringArray("hd-path")
		if err != nil {
			return err
		}
		seedB64 := strings.TrimSpace(cmd.Flag("seed").Value.String())
		kz, err := keys.EncodeSomeKeysFromSeed(hdPaths, pHrp, seedB64)
		if err != nil {
			return err
		}
		return util.Display(kz, cmd, os.Stdout)
	},
}

func init() {
	addFlags(CmdEncode, flagFormat, flagHDPath, flagHRP, flagSeed)
}
