package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/yaml.v3"
)

func addFlags(cmd *cobra.Command, opts ...func(cmd *cobra.Command)) {
	for _, fn := range opts {
		fn(cmd)
	}
}

var flagHRP = func(cmd *cobra.Command) {
	cmd.PersistentFlags().String("hrp", "", "Human readable prefix")
	_ = cmd.MarkFlagRequired("hrp")
}

var flagFormat = func(cmd *cobra.Command) {
	cmd.PersistentFlags().String("format", "", "The format out output the keys [json|yaml|plain]")
}

var flagHDPath = func(cmd *cobra.Command) {
	cmd.PersistentFlags().StringArray("hd-path", []string{}, "The bip44 hd path used to derive the extended Key")
}

var flagSeed = func(cmd *cobra.Command) {
	cmd.PersistentFlags().String("seed", "", "The base64 url encoded seed to use for the key derivation")
}

var flagLAddr = func(cmd *cobra.Command) {
	cmd.PersistentFlags().String("laddr", "0.0.0.0:9000", "The address:port to listen on")
}

var flagMnemonic = func(cmd *cobra.Command) { //nolint
	cmd.PersistentFlags().String("mnemonic", "", "The mnemonic to use to generate the seed")
}

func formatize(format string) (Formatter, error) {
	switch format {
	case "json":
		return json.Marshal, nil
	case "yaml":
	case "":
		return yaml.Marshal, nil
	default:
		return nil, fmt.Errorf("invalid format %s", format)
	}
	return nil, nil
}

func envOrSecret(name string) (string, error) {
	var value string
	if m, ok := os.LookupEnv(strings.ToUpper(name)); !ok {
		fmt.Printf("%s: ", name)
		bz, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			return "", err
		}
		value = string(bz)
		fmt.Printf("\n")
	} else {
		value = m
	}
	return value, nil
}

type Formatter = func(data interface{}) ([]byte, error)
