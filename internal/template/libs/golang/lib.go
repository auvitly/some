package golang

var Lib = func() any { return &lib{} }

type lib struct{}

func (*lib) AST() any { return &ast{} }
