package golang

import (
	goast "go/ast"
)

type inspectDecl struct{}

var inspecDeclImpl = new(inspectDecl)

func (*inspectDecl) Gens(v any) ([]*goast.GenDecl, error) {
	return any2node[*goast.GenDecl](v)
}

func (*inspectDecl) Gen() any {
	return &inspectDeclGen{}
}

func (*inspectDecl) Funcs(v any) ([]*goast.FuncDecl, error) {
	return any2node[*goast.FuncDecl](v)
}

func (*inspectDecl) Func() any {
	return &inspectDeclFunc{}
}

func (*inspectDecl) Bads(v any) ([]*goast.BadDecl, error) {
	return any2node[*goast.BadDecl](v)
}
