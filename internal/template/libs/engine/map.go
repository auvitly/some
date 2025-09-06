package engine

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/auvitly/gopher/internal/utils"
)

func Map(v ...any) (any, error) {
	switch {
	case len(v) == 0:
		return make(map[string]any), nil
	case len(v) == 1:
		var m = make(map[string]any)

		switch value := v[0].(type) {
		case []byte:
			err := json.Unmarshal(value, &m)
			if err != nil {
				return nil, err
			}

			return m, nil
		case string:
			err := json.Unmarshal([]byte(value), &m)
			if err != nil {
				return nil, err
			}

			return m, nil
		default:
			return nil, fmt.Errorf("cast to string failed: unsupported type %T", value)
		}
	default:
		return nil, fmt.Errorf("to many args")
	}
}

func Set(key string, value any, m any) (any, error) {
	var rv = reflect.ValueOf(m)

	if utils.Dereference(rv).Kind() != reflect.Map {
		return nil, fmt.Errorf("invalid type, must be map")
	}

	utils.Dereference(rv).SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))

	return rv.Interface(), nil
}

func Get(key string, m any) (any, error) {
	var rv = reflect.ValueOf(m)

	switch impl := m.(type) {
	case reflect.StructTag:
		return impl.Get(key), nil
	default:
		if utils.Dereference(rv).Kind() != reflect.Map {
			return nil, fmt.Errorf("invalid type, must be map")
		}

		return utils.Dereference(rv).MapIndex(reflect.ValueOf(key)).Interface(), nil
	}
}
