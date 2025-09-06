package maps

var Lib = func() any { return (*lib)(nil) }

type lib struct{}

func (*lib) Make() any { return (*makeLib)(nil) }
