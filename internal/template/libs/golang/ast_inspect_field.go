package golang

import (
	goast "go/ast"
	"reflect"
	"strings"
)

type inspectField struct{}

func (*inspectField) Names(v any) ([]*goast.Ident, error) {
	fields, err := any2node[*goast.Field](v)
	if err != nil {
		return nil, err
	}

	var names []*goast.Ident

	for _, field := range fields {
		names = append(names, field.Names...)
	}

	return names, nil
}

func (*inspectField) Types(v any) ([]goast.Expr, error) {
	fields, err := any2node[*goast.Field](v)
	if err != nil {
		return nil, err
	}

	var types []goast.Expr

	for _, field := range fields {
		types = append(types, field.Type)
	}

	return types, nil
}

// Tags.

func (*inspectField) Tags(v any) ([]*goast.BasicLit, error) {
	fields, err := any2node[*goast.Field](v)
	if err != nil {
		return nil, err
	}

	var basicLits []*goast.BasicLit

	for _, field := range fields {
		basicLits = append(basicLits, field.Tag)
	}

	return basicLits, nil
}

type inspectFieldTag struct{}

func (*inspectField) Tag() any { return (*inspectFieldTag)(nil) }

func (*inspectFieldTag) Values(v any) ([]reflect.StructTag, error) {
	fields, err := any2node[*goast.Field](v)
	if err != nil {
		return nil, err
	}

	var tags []reflect.StructTag

	for _, field := range fields {
		if field.Tag != nil {
			tags = append(tags, reflect.StructTag(strings.Trim(field.Tag.Value, "`")))
		} else {
			tags = append(tags, "")
		}
	}

	return tags, nil
}

// Docs.

func (*inspectField) Docs(v any) ([]*goast.CommentGroup, error) {
	fields, err := any2node[*goast.Field](v)
	if err != nil {
		return nil, err
	}

	var basicLits []*goast.CommentGroup

	for _, field := range fields {
		basicLits = append(basicLits, field.Doc)
	}

	return basicLits, nil
}

func (*inspectField) Comments(v any) ([]*goast.CommentGroup, error) {
	fields, err := any2node[*goast.Field](v)
	if err != nil {
		return nil, err
	}

	var basicLits []*goast.CommentGroup

	for _, field := range fields {
		basicLits = append(basicLits, field.Comment)
	}

	return basicLits, nil
}
