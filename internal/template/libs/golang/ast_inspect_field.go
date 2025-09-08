package golang

import (
	goast "go/ast"
	"strconv"
	"strings"

	"github.com/auvitly/gopher/internal/template/libs/golang/internal"
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

func (*inspectFieldTag) Values(v any) ([]internal.Tags, error) {
	fields, err := any2node[*goast.Field](v)
	if err != nil {
		return nil, err
	}

	var tags []internal.Tags

	for _, field := range fields {
		if field.Tag != nil {
			tags = append(tags, parseStructTag(strings.Trim(field.Tag.Value, "`")))
		} else {
			tags = append(tags, nil)
		}
	}

	return tags, nil
}

func parseStructTag(tag string) internal.Tags {
	var tags = make(internal.Tags)

	for tag != "" {
		i := 0

		for i < len(tag) && tag[i] == ' ' {
			i++
		}

		tag = tag[i:]

		if tag == "" {
			break
		}

		i = 0

		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}

		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}

		name := string(tag[:i])
		tag = tag[i+1:]

		i = 1

		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}

		if i >= len(tag) {
			break
		}

		qvalue := string(tag[:i+1])
		tag = tag[i+1:]

		value, err := strconv.Unquote(qvalue)
		if err != nil {
			break
		}

		tags[name] = value
	}

	return tags
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
