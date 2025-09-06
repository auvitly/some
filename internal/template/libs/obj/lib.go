package obj

import "fmt"

var Lib = func() any { return (*lib)(nil) }

type lib struct{}

func (*lib) Item() Object { return Object{} }
func (*lib) List() List   { return List{} }

func (*lib) Set(field string, value any, obj any) (any, error) {
	switch impl := obj.(type) {
	case List:
		return impl.Set(field, value)
	case Object:
		return impl.Set(field, value)
	default:
		return nil, fmt.Errorf("unsupported object type %T", impl)
	}
}

func (*lib) Get(field string, obj any) (any, error) {
	switch impl := obj.(type) {
	case List:
		return impl.Get(field), nil
	case Object:
		return impl.Get(field), nil
	default:
		return nil, fmt.Errorf("unsupported object type %T", impl)
	}
}

func (*lib) Del(field string, obj any) (any, error) {
	switch impl := obj.(type) {
	case List:
		return impl.Del(field), nil
	case Object:
		return impl.Del(field), nil
	default:
		return nil, fmt.Errorf("unsupported object type %T", impl)
	}
}
