package golang

import goast "go/ast"

type inspect struct{}

func (*inspect) Decl() any     { return new(inspectDecl) }
func (*inspect) Spec() any     { return new(inspectSpec) }
func (*inspect) Stmt() any     { return new(inspectStmt) }
func (*inspect) Field() any    { return new(inspectField) }
func (*inspect) Expr() any     { return new(inspectExpr) }
func (*inspect) File() any     { return new(inspectFile) }
func (*inspect) Package() any  { return new(inspectPackage) }
func (*inspect) Position() any { return new(inspectPosition) }

func (*inspect) Decls(v any) ([]goast.Decl, error)    { return any2node[goast.Decl](v) }
func (*inspect) Specs(v any) ([]goast.Spec, error)    { return any2node[goast.Spec](v) }
func (*inspect) Stmts(v any) ([]goast.Stmt, error)    { return any2node[goast.Stmt](v) }
func (*inspect) Fields(v any) ([]*goast.Field, error) { return any2node[*goast.Field](v) }
func (*inspect) Exprs(v any) ([]goast.Expr, error)    { return any2node[goast.Expr](v) }
func (*inspect) Files(v any) ([]*goast.File, error)   { return any2node[*goast.File](v) }
