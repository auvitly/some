package engine

func Void(...any) any { return "" }

func Coalesce(v ...any) (any, error) {
	switch len(v) {
	case 0:
		return nil, nil
	case 1:
		return v[0], nil
	default:
		for _, item := range append([]any{v[len(v)-1]}, v[:len(v)-1]...) {
			if item != nil {
				return item, nil
			}
		}
	}

	return nil, nil
}
