package reflect

import (
	"fmt"
	"reflect"
)

var LibType = func() any { return new(Type) }

type Type struct{}

func (*Type) Of(v any) string {
	return reflect.TypeOf(v).String()
}

func (*Type) Is(s string, v any) bool {
	return reflect.TypeOf(v).String() == s
}

func (*Type) NotIs(s string, v any) bool {
	return reflect.TypeOf(v).String() != s
}

func (t *Type) MustIs(s string, v any) (bool, error) {
	if !t.Is(s, v) {
		return false, fmt.Errorf("value must be %s, but %T", s, v)
	}

	return true, nil
}

func (t *Type) MustNotIs(s string, v any) (bool, error) {
	if !t.NotIs(s, v) {
		return false, fmt.Errorf("value must not be %s, but is", s)
	}

	return true, nil
}
