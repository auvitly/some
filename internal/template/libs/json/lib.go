package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Lib = func() any { return funcs{} }

type funcs struct{}

func (funcs) Marshal(v ...any) ([]byte, error) {
	switch len(v) {
	case 1:
		return json.Marshal(v[0])
	case 2:
		indent, ok := v[0].(string)
		if !ok {
			return nil, fmt.Errorf("indent must be string")
		}

		return json.MarshalIndent(v[1], "", indent)
	case 3:
		prefix, ok := v[0].(string)
		if !ok {
			return nil, fmt.Errorf("prefix must be string")
		}

		indent, ok := v[1].(string)
		if !ok {
			return nil, fmt.Errorf("indent must be string")
		}

		return json.MarshalIndent(v[2], prefix, indent)
	default:
		return nil, fmt.Errorf("incorrect arguments")
	}
}

func (funcs) Unmarshal(raw []byte) (any, error) {
	var parsed any

	err := json.Unmarshal(raw, &parsed)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func (funcs) Compression(raw []byte) ([]byte, error) {
	if !json.Valid(raw) {
		return nil, fmt.Errorf("is not a json value")
	}

	var v any

	err := json.Unmarshal(raw, &v)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (funcs) Read(name string) (any, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var v any

	err = json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
