package engine

func Doc(a any) any {
	if impl, ok := a.(interface{ Doc() any }); ok {
		return impl.Doc()
	}

	return "DOCUMENTATION NOT FOUND"
}
