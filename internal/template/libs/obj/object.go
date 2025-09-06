package obj

import (
	"fmt"
	"reflect"
)

type List []Object

type Object map[string]any

func (l List) Append(o Object) (List, error) {
	return append(l, o), nil
}

func (l List) Set(field string, value any) (List, error) {
	var rv = reflect.ValueOf(value)
	switch {
	case rv.Kind() != reflect.Slice:
		return nil, fmt.Errorf("value must be a slice")
	case rv.Len() != len(l):
		if len(l) != 0 {
			return nil, fmt.Errorf(
				"for field \"%s\" not match length %d/%d for set field",
				field, rv.Len(), len(l),
			)
		}

		for i := range rv.Len() {
			l = append(l, map[string]any{
				field: rv.Index(i).Interface(),
			})
		}

		return l, nil
	default:
		for i := range len(l) {
			_, err := l[i].Set(field, rv.Index(i).Interface())
			if err != nil {
				return nil, err
			}
		}

		return l, nil
	}
}

func (l List) Get(field string) any {
	var values []any

	for _, item := range l {
		values = append(values, item.Get(field))
	}

	return values
}

func (o Object) Set(field string, value any) (Object, error) {
	o[field] = value

	return o, nil
}

func (o Object) Get(field string) any {
	return o[field]
}

func (o List) Del(field string) List {
	for _, item := range o {
		item.Del(field)
	}

	return o
}

func (o Object) Del(field string) Object {
	delete(o, field)

	return o
}
