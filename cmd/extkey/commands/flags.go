package commands

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"syscall"
)

func addFlags(cmd *cobra.Command, opts ...func(cmd *cobra.Command)) {
	for _, fn := range opts {
		fn(cmd)
	}
}

var format string
var hdPath string
var hrp string
var seed string
var laddr string

var flagHRP = func(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&hrp, "hrp", "", "Human readable prefix")
	_ = cmd.MarkFlagRequired("hrp")
}

var flagFormat = func(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&format, "format", "", "The format out output the keys [json|yaml|plain]")
}

var flagHDPath = func(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&hdPath, "hd-path", "", "The bip44 hd path used to derive the extended Key")
}

var flagSeed = func(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&seed, "seed", "", "The base64 url encoded seed to use for the key derivation")
}

var flagLAddr = func(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&laddr, "laddr", "0.0.0.0:9000", "The address:port to listen on")
}

func formatize(format string) (Formatter, error) {
	if format == "json" {
		return func(data interface{}) ([]byte, error) {
			return json.Marshal(data)
		}, nil
	} else if format == "yaml" || format == "" {
		return func(data interface{}) ([]byte, error) {
			return yaml.Marshal(data)
		}, nil
	} else {
		return nil, fmt.Errorf("invalid format %s", format)
	}
}

func envOrPrompt(name string) (string, error) {
	var value string
	if m, ok := os.LookupEnv(strings.ToUpper(name)); !ok {
		fmt.Printf("%s: ", name)
		_, err := fmt.Scanf("%s", &value)
		if err != nil {
			return "", err
		}
	} else {
		hrp = m
	}
	return value, nil
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
