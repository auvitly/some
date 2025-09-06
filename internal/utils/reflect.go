package utils

import "reflect"

func MustSliceTo[T any, F any](from []F) []T {
	var to = make([]T, 0, len(from))

	for _, item := range from {
		to = append(to, (any(item)).(T))
	}

	return to
}

func Dereference(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	return v
}
