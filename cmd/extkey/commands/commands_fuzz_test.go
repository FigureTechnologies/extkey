package commands

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Golang doesn't support multi fuzz test cases. Aggregated in order to save time fuzzing
func FuzzCommands(f *testing.F) {
	f.Fuzz(func(t *testing.T, hrp string, path string, seed string, xkey string) {
		paths := []string{path}
		var err error
		var formatter Formatter
		formatter, err = formatize(strings.TrimSpace("json"))
		if err != nil {
			t.Error(err)
		} else {
			// Fuzz test generate and encode functions with fuzzy paths input. These should both panic most of the time with random input.
			if path != "m" {
				assert.Panics(t, func() { generate(hrp, "test", paths, formatter, os.Stdout) }, fmt.Sprintf("Test generate fuzz path did not panic. Path: %s", path))
				assert.Panics(t, func() { encode(paths, formatter, os.Stdout, hrp, "test") }, fmt.Sprintf("Test encode fuzz path did not panic. Path: %s", path))
			}

			// Fuzz test generate and encode functions with fuzzy seed. These should usually error with random inputs.
			if seed != "" && seed != "\r" && seed != "\n" && len(seed)%4 != 0 {
				paths := []string{"m/44'/505'/0'/0'/0"}
				generr := generate(hrp, seed, paths, formatter, os.Stdout)
				if generr == nil {
					t.Error(fmt.Sprintf("Test generate fuzz seed did not error. Seed: %s", seed))
				}
				encodeErr := encode(paths, formatter, os.Stdout, hrp, seed)
				if encodeErr == nil {
					t.Error(fmt.Sprintf("Test encode fuzz seed did not error. Seed: %s", seed))
				}
			}

			// Test fuzzy decode with random input.
			if xkey != "" {
				decodeErr := decode(hrp, xkey, os.Stdout, formatter)
				if decodeErr == nil {
					t.Error(fmt.Sprintf("Test decode fuzz xkey did not error. xkey: %s", xkey))
				}
			}
		}
	})
}
