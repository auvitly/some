package engine

import (
	"fmt"
)

func Line(v ...any) (any, error)  { return wrap("\n", v) }
func Tab(v ...any) (any, error)   { return wrap("\t", v) }
func Space(v ...any) (any, error) { return wrap(" ", v) }
func Dot(v ...any) (any, error)   { return wrap(".", v) }

func wrap(token string, v []any) (any, error) {
	switch len(v) {
	case 0:
		return token, nil
	case 1:
		switch value := v[0].(type) {
		case string:
			return fmt.Sprintf("%s%s", value, token), nil
		case []byte:
			return fmt.Sprintf("%s%s", string(value), token), nil
		default:
			return fmt.Sprintf("%v%s", value, token), nil
		}
	default:
		return nil, fmt.Errorf("too many args")
	}
}
