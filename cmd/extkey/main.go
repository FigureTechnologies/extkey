package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/enigmampc/btcutil/bech32"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ssh/terminal"
	"hash"
	"os"
	"golang.org/x/crypto/ripemd160"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) > 1 {
		decode(os.Args[1])
	} else {
		encode()
	}
}

func decode(xkey string) {
	var hrp string
	if h, ok := os.LookupEnv("HRP"); !ok {
		fmt.Printf("HRP: ")
		_, err := fmt.Scanf("%s", &hrp)
		if err != nil {
			panic(err)
		}
	} else {
		hrp = h
	}

	key, err := bip32.B58Deserialize(xkey)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Printf("Key Private: %s\n", key.B58Serialize())
	fmt.Printf("Key Public : %s\n", key.PublicKey().B58Serialize())
	fmt.Printf("ChainCode  : %X\n", key.ChainCode)
	fmt.Println()
	fmt.Printf("Depth      : %d (%s)\n", key.Depth, depthString(key.Depth))
	fmt.Printf("Address    : %s\n", toAddress(hrp, key))
}

func depthString(depth byte) string {
	path := []string{
		"m", "44", "coin", "account", "change", "index",
	}

	for i, _ := range path {
		if i == int(depth) {
			path[i] = "*"+path[i]
		}
	}
	return strings.Join(path, "/")
}

func encode() {
	var err error
	var mnemonic []byte
	if m, ok := os.LookupEnv("MNEMONIC"); !ok {
		fmt.Printf("Mnemonic: ")
		mnemonic, err = terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n\n")
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
		fmt.Printf("\n\n")
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
	fmt.Printf("Address: %s\n", toAddress(hrp, rootKey.PublicKey()))

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
	if key.IsPrivate {
		key = key.PublicKey()
	}

	addr, err := toAddressErr(hrp, key.Key)
	if err != nil {
		panic(err)
	}
	return addr
}

// toAddressErr converts from a base64 encoded byte string to base32 encoded byte string and then to bech32.
func toAddressErr(hrp string, data []byte) (string, error) {
	bz := Hash160(data)
	converted, err := bech32.ConvertBits(bz, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("encoding bech32 failed: %w", err)
	}
	return bech32.Encode(hrp, converted)
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

// Hash160 calculates the hash ripemd160(sha256(b)).
func Hash160(buf []byte) []byte {
	return calcHash(calcHash(buf, sha256.New()), ripemd160.New())
}

func calcHash(buf []byte, hasher hash.Hash) []byte {
	_, _ = hasher.Write(buf)
	return hasher.Sum(nil)
}