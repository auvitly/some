package golang

import (
	goast "go/ast"
)

type inspectExpr struct{}

// Ident.

func (*inspectExpr) Idents(v any) ([]*goast.Ident, error) {
	return any2node[*goast.Ident](v)
}

type inspectExprIdent struct{}

func (*inspectExpr) Ident() any { return &inspectExprIdent{} }

func (*inspectExprIdent) Names(v any) ([]any, error) {
	exprs, err := any2node[*goast.Ident](v)
	if err != nil {
		return nil, err
	}

	var list = make([]any, 0, len(exprs))

	for _, expr := range exprs {
		list = append(list, expr.Name)
	}

	return list, nil
}

func (*inspectExpr) Ellipsises(v any) ([]*goast.Ellipsis, error) {
	return any2node[*goast.Ellipsis](v)
}

// BasicLit.

func (*inspectExpr) BasicLits(v any) ([]*goast.BasicLit, error) {
	return any2node[*goast.BasicLit](v)
}

type inspectExprBasicLit struct{}

func (*inspectExpr) BasicLit() any { return &inspectExprBasicLit{} }

func (*inspectExprBasicLit) Values(v any) ([]any, error) {
	exprs, err := any2node[*goast.BasicLit](v)
	if err != nil {
		return nil, err
	}

	var list = make([]any, 0, len(exprs))

	for _, expr := range exprs {
		list = append(list, expr.Value)
	}

	return list, nil
}

// FuncLit.

func (*inspectExpr) FuncLits(v any) ([]*goast.FuncLit, error) {
	return any2node[*goast.FuncLit](v)
}

func (*inspectExpr) CompositeLits(v any) ([]*goast.CompositeLit, error) {
	return any2node[*goast.CompositeLit](v)
}

func (*inspectExpr) Parens(v any) ([]*goast.ParenExpr, error) {
	return any2node[*goast.ParenExpr](v)
}

func (*inspectExpr) Selectors(v any) ([]*goast.SelectorExpr, error) {
	return any2node[*goast.SelectorExpr](v)
}

func (*inspectExpr) Indexes(v any) ([]*goast.IndexExpr, error) {
	return any2node[*goast.IndexExpr](v)
}

func (*inspectExpr) IndexLists(v any) ([]*goast.IndexListExpr, error) {
	return any2node[*goast.IndexListExpr](v)
}

func (*inspectExpr) Slices(v any) ([]*goast.SliceExpr, error) {
	return any2node[*goast.SliceExpr](v)
}

func (*inspectExpr) TypeAsserts(v any) ([]*goast.TypeAssertExpr, error) {
	return any2node[*goast.TypeAssertExpr](v)
}

func (*inspectExpr) Calls(v any) ([]*goast.CallExpr, error) {
	return any2node[*goast.CallExpr](v)
}

func (*inspectExpr) Stars(v any) ([]*goast.StarExpr, error) {
	return any2node[*goast.StarExpr](v)
}

func (*inspectExpr) Unarys(v any) ([]*goast.UnaryExpr, error) {
	return any2node[*goast.UnaryExpr](v)
}

func (*inspectExpr) Binarys(v any) ([]*goast.BinaryExpr, error) {
	return any2node[*goast.BinaryExpr](v)
}

func (*inspectExpr) KeyValues(v any) ([]*goast.KeyValueExpr, error) {
	return any2node[*goast.KeyValueExpr](v)
}

func (*inspectExpr) ArrayTypes(v any) ([]*goast.ArrayType, error) {
	return any2node[*goast.ArrayType](v)
}

// StructType.

func (*inspectExpr) StructTypes(v any) ([]*goast.StructType, error) {
	return any2node[*goast.StructType](v)
}

type inspectExprStructType struct{}

func (*inspectExpr) StructType() any { return &inspectExprStructType{} }

func (*inspectExprStructType) FieldLists(v any) ([]*goast.FieldList, error) {
	list, err := any2node[*goast.StructType](v)
	if err != nil {
		return nil, err
	}

	var fieldLists []*goast.FieldList

	for _, item := range list {
		fieldLists = append(fieldLists, item.Fields)
	}

	return fieldLists, nil
}

func (*inspectExprStructType) Incomplete(v any) ([]*goast.StructType, error) {
	list, err := any2node[*goast.StructType](v)
	if err != nil {
		return nil, err
	}

	var results []*goast.StructType

	for _, item := range list {
		if item.Incomplete {
			results = append(results, item)
		}

	}

	return results, nil
}

func (*inspectExprStructType) Complete(v any) ([]*goast.StructType, error) {
	list, err := any2node[*goast.StructType](v)
	if err != nil {
		return nil, err
	}

	var results []*goast.StructType

	for _, item := range list {
		if !item.Incomplete {
			results = append(results, item)
		}

	}

	return results, nil
}

// FuncType.

func (*inspectExpr) FuncTypes(v any) ([]*goast.FuncType, error) {
	return any2node[*goast.FuncType](v)
}

func (*inspectExpr) InterfaceTypes(v any) ([]*goast.InterfaceType, error) {
	return any2node[*goast.InterfaceType](v)
}

func (*inspectExpr) MapTypes(v any) ([]*goast.MapType, error) {
	return any2node[*goast.MapType](v)
}

func (*inspectExpr) ChanTypes(v any) ([]*goast.ChanType, error) {
	return any2node[*goast.ChanType](v)
}
