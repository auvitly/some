package golang

import (
	goast "go/ast"
)

type inspectStmt struct{}

func (*inspectStmt) Blocks(v any) ([]*goast.BlockStmt, error) {
	return any2node[*goast.BlockStmt](v)
}

func (*inspectStmt) Assigns(v any) ([]*goast.AssignStmt, error) {
	return any2node[*goast.AssignStmt](v)
}

func (*inspectStmt) Selects(v any) ([]*goast.SelectStmt, error) {
	return any2node[*goast.SelectStmt](v)
}

func (*inspectStmt) Fors(v any) ([]*goast.ForStmt, error) {
	return any2node[*goast.ForStmt](v)
}

func (*inspectStmt) Branchs(v any) ([]*goast.BranchStmt, error) {
	return any2node[*goast.BranchStmt](v)
}

func (*inspectStmt) Defers(v any) ([]*goast.DeferStmt, error) {
	return any2node[*goast.DeferStmt](v)
}

func (*inspectStmt) Exprs(v any) ([]*goast.ExprStmt, error) {
	return any2node[*goast.ExprStmt](v)
}

func (*inspectStmt) Decls(v any) ([]*goast.DeclStmt, error) {
	return any2node[*goast.DeclStmt](v)
}

func (*inspectStmt) Gos(v any) ([]*goast.GoStmt, error) {
	return any2node[*goast.GoStmt](v)
}

func (*inspectStmt) Ifs(v any) ([]*goast.IfStmt, error) {
	return any2node[*goast.IfStmt](v)
}

func (*inspectStmt) Ranges(v any) ([]*goast.RangeStmt, error) {
	return any2node[*goast.RangeStmt](v)
}

func (*inspectStmt) Sends(v any) ([]*goast.SendStmt, error) {
	return any2node[*goast.SendStmt](v)
}

func (*inspectStmt) TypeSwitchs(v any) ([]*goast.TypeSwitchStmt, error) {
	return any2node[*goast.TypeSwitchStmt](v)
}

func (*inspectStmt) Labeleds(v any) ([]*goast.LabeledStmt, error) {
	return any2node[*goast.LabeledStmt](v)
}

func (*inspectStmt) Returns(v any) ([]*goast.ReturnStmt, error) {
	return any2node[*goast.ReturnStmt](v)
}

func (*inspectStmt) Emptys(v any) ([]*goast.EmptyStmt, error) {
	return any2node[*goast.EmptyStmt](v)
}
