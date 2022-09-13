package main

import (
	"github.com/FigureTechnologies/extkey/cmd/extkey/commands"
)

func main() {
	if err := commands.CmdRoot.Execute(); err != nil {
		panic(err)
	}
}
