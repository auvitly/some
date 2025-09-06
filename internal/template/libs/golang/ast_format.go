package golang

import (
	"bytes"
	"fmt"
	goast "go/ast"
	goformat "go/format"
	"go/printer"
	"go/token"
	"reflect"
	"strings"

	"github.com/auvitly/gopher/internal/template/libs/golang/internal"
	"golang.org/x/tools/go/packages"
)

type format struct{}

func (*format) Node(values ...any) (any, error) {
	if len(values) == 0 {
		return nil, fmt.Errorf("values must be not empty")
	}

	var (
		set   *token.FileSet
		nodes []any
	)

	for _, value := range values {
		switch impl := value.(type) {
		case *token.FileSet:
			if set == nil {
				set = impl

				continue
			}

			return nil, fmt.Errorf("*token.FileSet should not be or should be in a single representation")
		default:
			nodes = append(nodes, value)
		}
	}

	if set == nil {
		set = token.NewFileSet()
	}

	var outs []string

	for _, node := range nodes {
		out, err := any2string(set, node)
		if err != nil {
			return nil, err
		}

		outs = append(outs, out)
	}

	return strings.Join(outs, "\n"), nil
}

func any2string(fset *token.FileSet, node any) (string, error) {
	var rv = reflect.ValueOf(node)

	var list []any

	switch rv.Kind() {
	case reflect.Slice:
		for i := range rv.Len() {
			list = append(list, rv.Index(i).Interface())
		}
	default:
		list = append(list, rv.Interface())
	}

	var buf = bytes.NewBuffer(nil)

	for i, item := range list {
		switch impl := item.(type) {
		case *internal.File:
			err := goformat.Node(buf, impl.Fset, impl.Syntax)
			if err != nil {
				return "", err
			}
		case *internal.Project:
			for _, pkg := range impl.Packages {
				err := goformat.Node(buf, pkg.Fset, impl)
				if err != nil {
					return "", err
				}
			}
		case *packages.Package:
			err := goformat.Node(buf, impl.Fset, impl.Syntax)
			if err != nil {
				return "", err
			}
		case *goast.Field:
			field2string(buf, fset, impl)
		case *goast.FieldList:
			for i, item := range impl.List {
				field2string(buf, fset, item)

				if i != len(impl.List)-1 {
					buf.WriteRune('\n')
				}
			}
		case goast.Node:
			err := goformat.Node(buf, fset, impl)
			if err != nil {
				return "", err
			}
		case string:
			buf.WriteString(impl)
		case reflect.StructTag:
			if len(string(impl)) != 0 {
				buf.WriteString(fmt.Sprintf("`%s`", string(impl)))
			}
		default:
			return "", fmt.Errorf("unsupported type: %T", impl)
		}

		if len(list) != 1 && i < len(list)-1 {
			buf.WriteString("\n")
		}
	}

	return buf.String(), nil
}

func field2string(buf *bytes.Buffer, fset *token.FileSet, field *goast.Field) string {
	if len(field.Names) > 0 {
		for i, name := range field.Names {
			if i > 0 {
				buf.WriteString(", ")
			}

			buf.WriteString(name.Name)
		}

		buf.WriteString(" ")
	}

	if field.Type != nil {
		var typeBuf bytes.Buffer

		printer.Fprint(&typeBuf, fset, field.Type)

		buf.WriteString(typeBuf.String())
	}

	if field.Tag != nil {
		buf.WriteString(" ")
		buf.WriteString(field.Tag.Value)
	}

	return buf.String()
}
