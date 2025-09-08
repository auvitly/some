package must

import (
	"fmt"
	"reflect"
)

var Lib = func() any { return &lib{} }

type lib struct{}

func (*lib) Nil(v any) (any, error) {
	if v != nil {
		return nil, fmt.Errorf("value must be nil")
	}

	return v, nil
}

func (*lib) NotNil(v any) (any, error) {
	if v == nil {
		return nil, fmt.Errorf("value must be not nil")
	}

	return v, nil
}

func (*lib) True(ok bool) (bool, error) {
	if !ok {
		return false, fmt.Errorf("condition must be true")
	}

	return true, nil
}

func (*lib) False(ok bool) (bool, error) {
	if ok {
		return false, fmt.Errorf("condition must be false")
	}

	return true, nil
}

func (*lib) Eq(x, y any) (any, error) {
	if !reflect.DeepEqual(x, y) {
		return nil, fmt.Errorf("values must be equal: %v!=%v", x, y)
	}

	return y, nil
}

func (*lib) NotEq(x, y any) (any, error) {
	if reflect.DeepEqual(x, y) {
		return nil, fmt.Errorf("values must be not equal: %v==%v", x, y)
	}

	return y, nil
}

func (*lib) Type(t string, v any) (any, error) {
	if t != reflect.TypeOf(v).String() {
		return nil, fmt.Errorf("type must be not equal: %s ~ %T", t, v)
	}

	return v, nil
}

func (*lib) TypeNot(t string, v any) (any, error) {
	if t == reflect.TypeOf(v).String() {
		return nil, fmt.Errorf("type must be not equal: %s", t)
	}

	return v, nil
}
