package engine

import (
	"fmt"
	"strconv"
)

func Int(v any) (int, error) {
	switch value := v.(type) {
	case int:
		return value, nil
	case int8:
		return int(value), nil
	case int16:
		return int(value), nil
	case int32:
		return int(value), nil
	case int64:
		return int(value), nil
	case uint:
		return int(value), nil
	case uint8:
		return int(value), nil
	case uint16:
		return int(value), nil
	case uint32:
		return int(value), nil
	case uint64:
		return int(value), nil
	case float32:
		return int(value), nil
	case float64:
		return int(value), nil
	case string:
		res, err := strconv.ParseInt(value, 10, 64)

		return int(res), err
	default:
		return 0, fmt.Errorf("cast to string failed: unsupported type %T", value)
	}
}
