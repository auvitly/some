package maps

import (
	"fmt"
	"reflect"
)

type makeLib struct{}

func (*makeLib) Zip(keys any, values any) (map[any]any, error) {
	var (
		kv = reflect.ValueOf(keys)
		vv = reflect.ValueOf(values)
	)

	if kv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("keys must be slice")
	}

	if vv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("values must be slice")
	}

	if kv.Len() != vv.Len() {
		return nil, fmt.Errorf("not match lenghts")
	}

	var m = make(map[any]any)

	for i := range kv.Len() {
		m[kv.Index(i).Interface()] = vv.Index(i).Interface()
	}

	return m, nil
}

func (*makeLib) FromKeys(keys []any) (map[any]any, error) {
	var m = make(map[any]any)

	for _, key := range keys {
		m[key] = nil
	}

	return m, nil
}
