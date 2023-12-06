package commands

import (
	"github.com/spf13/cobra"
)

var CmdRoot = &cobra.Command{
	Use: "extkey",
}

func init() {
	CmdRoot.AddCommand(CmdGenerate, CmdEncode, CmdServe)
}
