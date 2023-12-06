package util

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
)

func Display(v any, format string, w io.Writer) error {
	formatter, err := NewFormatter(format)
	if err != nil {
		return err
	}
	output, err := formatter(v)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%s\n", output)
	return nil
}

func NewFormatter(format string) (Formatter, error) {
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

type Formatter = func(data interface{}) ([]byte, error)
