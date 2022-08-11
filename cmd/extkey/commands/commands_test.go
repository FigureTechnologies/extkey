package commands

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)


func TestGenerate(t *testing.T) {
	paths := []string{"m/44'/505'/0'/0'/0"}

	r, w, _ := os.Pipe()
	os.Stdout = w

	expected := "xpub661MyMwAqRbcEgfM91SZpDwJMa4xs8Kanuno3s1NXfwvbQ93vryox2hZLZjeDm8cSGNbP8BbFy3w29udT2HeQdKfTbWQkekAtdnDakM1gVG"
	formatter, err := formatize(strings.TrimSpace("json"))
	if err != nil {
		t.Error(err)
	} else {
		genErr := generate("tp", "test", paths, formatter, os.Stdout)
		if genErr != nil {
			t.Error(genErr)
		}
		w.Close()

		out, _ := ioutil.ReadAll(r)

		genSeedlessErr := generate("tp", "", paths, formatter, os.Stdout)
		if genSeedlessErr != nil {
			t.Error(genSeedlessErr)
		}

		genFailErr := generate("tp", "testf", paths, formatter, os.Stdout)
		if genFailErr == nil {
			t.Error("Error expected but not thrown")
		}

		var genKey SomeKey
		jsonErr := json.Unmarshal(out, &genKey)
		if jsonErr != nil {
			t.Error(jsonErr)
		}

		if genKey.RootKey.PublicKey.Base58 != expected {
			t.Error("Unexpected pub key generated")
		}
	}
}

func TestEncode(t *testing.T) {
	paths := []string{"m/44'/1'/0'/420'"}

	expected := "xpub661MyMwAqRbcEgfM91SZpDwJMa4xs8Kanuno3s1NXfwvbQ93vryox2hZLZjeDm8cSGNbP8BbFy3w29udT2HeQdKfTbWQkekAtdnDakM1gVG"
	r, w, _ := os.Pipe()
	os.Stdout = w

	formatter, err := formatize(strings.TrimSpace("json"))
	if err != nil {
		t.Error(err)
	} else {
		encodeErr := encode(paths, formatter, os.Stdout, "tp", "test")
		if encodeErr != nil {
			t.Error(encodeErr)
		}
		w.Close()

		encodeFailErr := encode(paths, formatter, os.Stdout, "tp", "testf")
		if encodeFailErr == nil {
			t.Error("Error expected but not thrown")
		}

		out, _ := ioutil.ReadAll(r)

		var genKey SomeKey
		jsonErr := json.Unmarshal(out, &genKey)
		if jsonErr != nil {
			t.Error(jsonErr)
		}

		if genKey.RootKey.PublicKey.Base58 != expected {
			t.Error("Unexpected pub key generated")
		}
	}
}

func TestDecode(t *testing.T) {
	formatter, err := formatize(strings.TrimSpace("json"))
	r, w, _ := os.Pipe()
	os.Stdout = w

	expected := "0488B21E0000000000000000000E3F9A93AB0FCAB898F0BA7FA05B21262AC8C3D313E99D9EE7AEA7CDF864205B023EB1E5184C5398A591F72751FB0104A6F1C9817BF6744C6E24E3F1742AC33A128BC11343"
	if err != nil {
		t.Error(err)
	} else {
		decodeErr := decode("tp", "xpub661MyMwAqRbcEgfM91SZpDwJMa4xs8Kanuno3s1NXfwvbQ93vryox2hZLZjeDm8cSGNbP8BbFy3w29udT2HeQdKfTbWQkekAtdnDakM1gVG", os.Stdout, formatter)
		if decodeErr != nil {
			t.Error(decodeErr)
		}
		w.Close()

		out, _ := ioutil.ReadAll(r)

		var key decodedKeyInfo
		jsonErr := json.Unmarshal(out, &key)

		if jsonErr != nil {
			t.Error(jsonErr)
		}

		if key.XKey.RootKey.PublicKey.Bytes != expected {
			t.Error("Invalid decoding")
		}
	}
}

func TestCommandsCompatibility(t *testing.T) {
	paths := []string{"m/44'/505'/0'/0'/0"}

	r, w, _ := os.Pipe()
	os.Stdout = w

	formatter, err := formatize(strings.TrimSpace("json"))

	var genKey SomeKey
	var decodeKey decodedKeyInfo
	var encodeKey SomeKey

	if err != nil {
		t.Error(err)
	} else {
		genErr := generate("tp", "test", paths, formatter, os.Stdout)
		if genErr != nil {
			t.Error(genErr)
		}
		w.Close()
		genOut, _ := ioutil.ReadAll(r)

		jsonErr := json.Unmarshal(genOut, &genKey)
		if jsonErr != nil {
			t.Error(jsonErr)
		}

		r, w, _ := os.Pipe()
		os.Stdout = w

		decodeErr := decode("tp", genKey.RootKey.PublicKey.Base58, os.Stdout, formatter)
		if decodeErr != nil {
			t.Error(decodeErr)
		}
		w.Close()

		decodeOut, _ := ioutil.ReadAll(r)

		jsonErr = json.Unmarshal(decodeOut, &decodeKey)
		if jsonErr != nil {
			t.Error(jsonErr)
		}

		r, w, _ = os.Pipe()
		os.Stdout = w

		encodeErr := encode(paths, formatter, os.Stdout, "tp", "test")
		if encodeErr != nil {
			t.Error(encodeErr)
		}
		w.Close()

		encodeOut, _ := ioutil.ReadAll(r)

		jsonErr = json.Unmarshal(encodeOut, &encodeKey)
		if jsonErr != nil {
			t.Error(jsonErr)
		}

		if encodeKey.RootKey.Address != genKey.RootKey.Address {
			t.Error("Encoded key value isn't equal to the generate value which should be identical")
		}
		if encodeKey.RootKey.Address != decodeKey.XKey.RootKey.Address {
			t.Error("Decoded value doesn't match expected from encode function")
		}
		if decodeKey.XKey.RootKey.Address != genKey.RootKey.Address {
			t.Error("Decoded value doesn't match expected value from encode function")
		}
	}
}
