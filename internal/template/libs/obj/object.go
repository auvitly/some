package obj

import (
	"fmt"
	"reflect"
)

type List []Item

type Item map[string]any

func (l List) Append(o Item) (List, error) {
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
			if l[i] == nil {
				l[i] = make(Item)
			}

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

func (o Item) Set(field string, value any) (Item, error) {
	o[field] = value

	return o, nil
}

func (o Item) Get(field string) any {
	return o[field]
}

func (o List) Del(field string) List {
	for _, item := range o {
		item.Del(field)
	}

	return o
}

func (o Item) Del(field string) Item {
	delete(o, field)

	return o
}
