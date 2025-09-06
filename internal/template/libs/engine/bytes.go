package engine

import "fmt"

func Bytes(v any) ([]byte, error) {
	switch value := v.(type) {
	case []byte:
		return value, nil
	case string:
		return []byte(value), nil
	default:
		return nil, fmt.Errorf("cast to string failed: unsupported type %T", value)
	}
}
