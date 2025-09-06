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

func Impl(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Interface || v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	return v
}

func AnyTo[T any](v ...any) ([]T, error) {
	var (
		values = reflect.ValueOf(v)
		target = reflect.TypeOf(*new(T))
		casted = make([]T, 0, len(v))
	)

	for i := range values.Len() {
		var value = Impl(values.Index(i))

		switch {
		case target.Kind() == value.Kind():
			if impl, ok := value.Interface().(T); ok {
				casted = append(casted, impl)
			}
		case value.Kind() == reflect.Slice:
			for j := range value.Len() {
				items, err := AnyTo[T](value.Index(j).Interface())
				if err != nil {
					return nil, err
				}

				casted = append(casted, items...)
			}
		}
	}

	return casted, nil
}
