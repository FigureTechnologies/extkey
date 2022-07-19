package main

import (
	"github.com/provenance-io/extkey/cmd/extkey/commands"
)

func main() {
	if err := commands.CmdRoot.Execute(); err != nil {
		panic(err)
	}
}
