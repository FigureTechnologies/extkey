package commands

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/gogo/protobuf/proto"

	"github.com/FigureTechnologies/extkey/pkg/encryption/eckey"

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
		formatter, err := formatize(strings.TrimSpace(cmd.Flag("format").Value.String()))
		if err != nil {
			return err
		}
		return decodeProto(protoBz, os.Stdout, formatter, pub)
	},
}

func init() {
	addFlags(CmdDecode, flagHRP, flagFormat)

	cmdDecodeProto.PersistentFlags().Bool("pub", false, "Decode public version of the proto bytes")
	CmdDecode.AddCommand(cmdDecodeProto)
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

func decodeProto(bz []byte, w io.Writer, formatter Formatter, isPub bool) error {
	msg := eckey.Key{}
	if err := proto.Unmarshal(bz, &msg); err != nil {
		return err
	}

	var curve *btcec.KoblitzCurve
	switch msg.Curve {
	case eckey.KeyCurve_SECP256K1:
		curve = btcec.S256()
	default:
		panic("curve not found " + msg.Curve.String())
	}

	var pub *InnerKeyData
	var prv *InnerKeyData
	if isPub {
		pk, err := btcec.ParsePubKey(msg.KeyBytes, curve)
		if err != nil {
			return err
		}
		if !curve.IsOnCurve(pk.X, pk.Y) {
			return fmt.Errorf("invalid point for curve")
		}
		pub = NewInnerKeyDataFromBTCECPub(pk)
	} else {
		prvK, pubK := btcec.PrivKeyFromBytes(curve, msg.KeyBytes)
		if prvK == nil || pubK == nil {
			return fmt.Errorf("invalid private key")
		}
		if !curve.IsOnCurve(prvK.X, prvK.Y) {
			return fmt.Errorf("invalid point for private key")
		}
		if !curve.IsOnCurve(pubK.X, pubK.Y) {
			return fmt.Errorf("invalid point for public key")
		}
		prv, pub = NewInnerKeyDataFromBTCECPrv(prvK)
	}

	obj := decodedProtoInfo{
		PublicKey:  pub,
		PrivateKey: prv,
	}
	output, err := formatter(obj)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%s\n", output)
	return nil
}

type decodedProtoInfo struct {
	PrivateKey *InnerKeyData `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`
	PublicKey  *InnerKeyData `json:"publicKey" yaml:"publicKey"`
}
