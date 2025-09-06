package engine

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Slice(v ...any) ([]any, error) {
	switch {
	case len(v) == 0:
		return make([]any, 0), nil
	case len(v) == 1:
		var s = make([]any, 0)

		switch value := v[0].(type) {
		case []byte:
			err := json.Unmarshal(value, &s)
			if err != nil {
				return nil, err
			}

			return s, nil
		case string:
			err := json.Unmarshal([]byte(value), &s)
			if err != nil {
				return []any{value}, nil
			}

			return s, nil
		case []any:
			return value, nil
		default:
			return nil, fmt.Errorf("cast to string failed: unsupported type %T", value)
		}
	default:
		return v, nil
	}
}

func Unique(v any) ([]any, error) {
	var (
		rv     = reflect.ValueOf(v)
		unique []reflect.Value
	)

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, fmt.Errorf("must be slice or array: current %T", v)
	}

loop:
	for i := range rv.Len() {
		for _, item := range unique {
			if item.Equal(rv.Index(i)) {
				continue loop
			}
		}

		unique = append(unique, rv.Index(i))
	}

	var result = make([]any, 0, len(unique))

	for _, item := range unique {
		result = append(result, item.Interface())
	}

	return result, nil
}

func Contains(e any, s []any) bool {
	var (
		ve = reflect.ValueOf(e)
		vs = reflect.ValueOf(s)
	)

	for i := range vs.Len() {
		if ve.Equal(vs.Index(i)) {
			return true
		}
	}

	return false
}

func Append(elem any, slice []any) []any {
	return append(slice[:], elem)
}

func Index(index int, slice any) (any, error) {
	var rv = reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("invalid type, must be slice")
	}

	if rv.Len() < index {
		return nil, fmt.Errorf("out of range[%d] (len:%d)", index, rv.Len())
	}

	return rv.Index(index).Interface(), nil
}

func First(slice any) (any, error) {
	var rv = reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("invalid type, must be slice")
	}

	if rv.Len() < 1 {
		return nil, fmt.Errorf("out of range")
	}

	return rv.Index(0).Interface(), nil
}
