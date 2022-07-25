package commands

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/bech32"
	bip32 "github.com/tyler-smith/go-bip32"
	"golang.org/x/crypto/ripemd160"
)

func depthString(depth byte) string {
	path := []string{
		"m", "44", "coin", "account", "change", "index",
	}

	for i := range path {
		if i == int(depth) {
			path[i] = "*" + path[i]
		}
	}
	return strings.Join(path, "/")
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
