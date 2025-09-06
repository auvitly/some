package golang

import goast "go/ast"

type inspect struct{}

func (*inspect) Decl() any     { return (*inspectDecl)(nil) }
func (*inspect) Spec() any     { return (*inspectSpec)(nil) }
func (*inspect) Stmt() any     { return (*inspectStmt)(nil) }
func (*inspect) Field() any    { return (*inspectField)(nil) }
func (*inspect) Expr() any     { return (*inspectExpr)(nil) }
func (*inspect) File() any     { return (*inspectFile)(nil) }
func (*inspect) Package() any  { return (*inspectPackage)(nil) }
func (*inspect) Position() any { return (*inspectPosition)(nil) }

func (*inspect) Decls(v any) ([]goast.Decl, error)    { return any2node[goast.Decl](v) }
func (*inspect) Specs(v any) ([]goast.Spec, error)    { return any2node[goast.Spec](v) }
func (*inspect) Stmts(v any) ([]goast.Stmt, error)    { return any2node[goast.Stmt](v) }
func (*inspect) Fields(v any) ([]*goast.Field, error) { return any2node[*goast.Field](v) }
func (*inspect) Exprs(v any) ([]goast.Expr, error)    { return any2node[goast.Expr](v) }
func (*inspect) Files(v any) ([]*goast.File, error)   { return any2node[*goast.File](v) }
