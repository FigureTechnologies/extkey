package main

import (
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	var err error
	var mnemonic []byte
	if m, ok := os.LookupEnv("MNEMONIC"); !ok {
		fmt.Printf("Mnemonic: ")
		mnemonic, err = terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n")
	} else {
		mnemonic = []byte(m)
	}

	var passphrase []byte
	if p, ok := os.LookupEnv("PASSPHRASE"); !ok {
		fmt.Printf("Passphrase: ")
		passphrase, err = terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n")
	} else {
		passphrase = []byte(p)
	}

	var hrp string
	if h, ok := os.LookupEnv("HRP"); !ok {
		fmt.Printf("HRP: ")
		_, err = fmt.Scanf("%s", &hrp)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n")
	} else {
		hrp = h
	}


	var path string
	if r, ok := os.LookupEnv("HDPATH"); !ok {
		fmt.Printf("HDPath: ")
		_, err = fmt.Scanf("%s", &path)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n")
	} else {
		path = r
	}

	seed := bip39.NewSeed(string(mnemonic), string(passphrase))
	rootKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Printf("RootKey Private: %s\n", rootKey.B58Serialize())
	fmt.Printf("RootKey Public : %s\n", rootKey.PublicKey().B58Serialize())
	fmt.Printf("Address: %s\n", toAddress(hrp, rootKey))

	childKey := rootKey
	bip44Indexes, bip44Harden, err := parseBIP44(path)
	if err != nil {
		panic(err)
	}
	for i, childIndex := range bip44Indexes {
		if bip44Harden[i] {
			childIndex |= 0x80000000
		}
		childKey, err = childKey.NewChildKey(childIndex)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println()
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("ChildKey Private: %s\n", childKey.B58Serialize())
	fmt.Printf("ChildKey Public : %s\n", childKey.PublicKey().B58Serialize())
	fmt.Printf("Address: %s\n", toAddress(hrp, childKey))
}

func toAddress(hrp string, key *bip32.Key) string {
	addr, err := toAddressErr(hrp, key)
	if err != nil {
		panic(err)
	}
	return addr
}

func toAddressErr(hrp string, key *bip32.Key) (string, error) {
	pubkey := key.PublicKey().Key
	conv, err := bech32.ConvertBits(pubkey, 8, 5, true)
	if err != nil {
		return "", err
	}
	address, err := bech32.Encode(hrp, conv)
	if err != nil {
		return "", err
	}
	return address, nil
}

func parseBIP44(path string) ([]uint32, []bool, error) {
	split := strings.Split(path, "/")
	if split[0] != "m" {
		return nil, nil, fmt.Errorf("invalid bip44 path")
	}
	if len(split) > 6 {
		return nil, nil, fmt.Errorf("bip44 path too deep")
	}
	indexes := make([]uint32, len(split)-1)
	harden := make([]bool, len(split)-1)
	for i, element := range split[1:] {
		toParse := element
		if strings.HasSuffix(element, "'") || strings.HasSuffix(element, "H") {
			harden[i] = true
			toParse = element[:len(element)-1]
		}
		childIndex, err := strconv.ParseUint(toParse, 10, 32)
		if err != nil {
			return nil, nil, err
		}
		indexes[i] = uint32(childIndex)
	}
	return indexes, harden, nil
}