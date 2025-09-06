package math

import (
	"fmt"
	"reflect"
)

var Lib = func() any { return &lib{} }

type lib struct{}

func (*lib) Sum(a, b any) (any, error) {
	var (
		va = reflect.ValueOf(a)
		vb = reflect.ValueOf(b)
	)

	switch {
	case va.CanInt() && vb.CanInt():
		return int(va.Int() + vb.Int()), nil
	case va.CanUint() && vb.CanUint():
		return uint(va.Uint() + vb.Uint()), nil
	case va.CanFloat() && vb.CanFloat():
		return float64(va.Float() + vb.Float()), nil
	case va.CanInt() && vb.CanFloat():
		return float64(va.Int()) + vb.Float(), nil
	case va.CanFloat() && vb.CanInt():
		return float64(va.Float() + float64(vb.Int())), nil
	default:
		return nil, fmt.Errorf("unsupported types for sum %T and %T", a, b)
	}
}

func (*lib) Div(a, b any) float64 {
	return reflect.ValueOf(a).Float() / reflect.ValueOf(b).Float()
}

func (*lib) Mul(a, b any) float64 {
	return reflect.ValueOf(a).Float() / reflect.ValueOf(b).Float()
}

func (*lib) Inc(a int) (any, error) {
	var va = reflect.ValueOf(a)

	switch {
	case va.CanInt():
		return int(va.Int() + 1), nil
	case va.CanUint():
		return uint(va.Uint() + 1), nil
	case va.CanFloat():
		return float64(va.Float() + 1), nil
	default:
		return nil, fmt.Errorf("unsupported types for inc %T ", a)
	}
}

func (*lib) Dec(a int) (any, error) {
	var va = reflect.ValueOf(a)

	switch {
	case va.CanInt():
		return int(va.Int() - 1), nil
	case va.CanUint():
		return uint(va.Uint() - 1), nil
	case va.CanFloat():
		return float64(va.Float() - 1), nil
	default:
		return nil, fmt.Errorf("unsupported types for dec %T", a)
	}
}
