package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"reflect"
)

func String(v any) (string, error) {
	switch value := v.(type) {
	case interface{ String() (string, error) }:
		return value.String()
	case []byte:
		return string(value), nil
	case string:
		return value, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return fmt.Sprintf("%v", value), nil
	default:
		var rv = reflect.ValueOf(v)

		switch rv.Kind() {
		case reflect.Array, reflect.Slice, reflect.Map:
			raw, err := json.Marshal(v)
			if err != nil {
				break
			}

			return string(raw), nil
		}

		return fmt.Sprintf("%v", value), nil
	}
}

func nodes2String(nodes ...ast.Node) (string, error) {
	var buf = bytes.NewBuffer(nil)

	for _, node := range nodes {
		err := format.Node(buf, token.NewFileSet(), node)
		if err != nil {
			return "", err
		}
	}

	return buf.String(), nil
}
