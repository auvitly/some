package obj

import "fmt"

var Lib = func() any { return (*lib)(nil) }

type lib struct{}

func (*lib) Item() Item { return Item{} }
func (*lib) List(v ...int) (List, error) {
	switch len(v) {
	case 0:
		return List{}, nil
	case 1:
		return make(List, v[0]), nil
	default:
		return List{}, fmt.Errorf("too many args")
	}
}

func (*lib) Set(field string, value any, obj any) (any, error) {
	switch impl := obj.(type) {
	case List:
		return impl.Set(field, value)
	case Item:
		return impl.Set(field, value)
	default:
		return nil, fmt.Errorf("unsupported object type %T", impl)
	}
}

func (*lib) Get(field string, obj any) (any, error) {
	switch impl := obj.(type) {
	case List:
		return impl.Get(field), nil
	case Item:
		return impl.Get(field), nil
	default:
		return nil, fmt.Errorf("unsupported object type %T", impl)
	}
}

func (*lib) Del(field string, obj any) (any, error) {
	switch impl := obj.(type) {
	case List:
		return impl.Del(field), nil
	case Item:
		return impl.Del(field), nil
	default:
		return nil, fmt.Errorf("unsupported object type %T", impl)
	}
}
