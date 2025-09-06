package engine

import (
	"fmt"
	"strconv"
)

func Bool(v any) (bool, error) {
	switch value := v.(type) {
	case int:
		return value != 0, nil
	case int8:
		return value != 0, nil
	case int16:
		return value != 0, nil
	case int32:
		return value != 0, nil
	case int64:
		return value != 0, nil
	case uint:
		return value != 0, nil
	case uint8:
		return value != 0, nil
	case uint16:
		return value != 0, nil
	case uint32:
		return value != 0, nil
	case uint64:
		return value != 0, nil
	case float32:
		return value != 0, nil
	case float64:
		return value != 0, nil
	case string:
		return strconv.ParseBool(value)
	case bool:
		return value, nil
	default:
		return false, fmt.Errorf("cast to string failed: unsupported type %T", value)
	}
}
