package golang

import (
	goast "go/ast"
	"reflect"

	"github.com/auvitly/gopher/internal/template/libs/golang/internal"
	"golang.org/x/tools/go/packages"
)

type ast struct{}

// Fullnames.

func (*ast) Format(v ...any) (any, error) { return new(format).Node(v...) }
func (*ast) Inspect() any                 { return new(inspect) }
func (*ast) Parse() any                   { return new(parse) }

// Shortcuts.

func (a *ast) F(v ...any) (any, error) { return a.Format(v...) }
func (a *ast) I() any                  { return a.Inspect() }
func (a *ast) P() any                  { return a.Parse() }

func any2node[N goast.Node](v any) ([]N, error) {
	var rv = reflect.ValueOf(v)

	var list []any

	switch rv.Kind() {
	case reflect.Slice:
		for i := range rv.Len() {
			list = append(list, rv.Index(i).Interface())
		}
	default:
		list = append(list, rv.Interface())
	}

	var results []N

	for _, item := range list {
		switch impl := item.(type) {
		case goast.Node:
			goast.Inspect(impl, astInspectNode(&results))
		case *packages.Package:
			for _, file := range impl.Syntax {
				goast.Inspect(file, astInspectNode(&results))
			}
		case *internal.File:
			goast.Inspect(impl.Syntax, astInspectNode(&results))
		case *internal.Project:
			for _, pkg := range impl.Packages {
				for _, file := range pkg.Syntax {
					goast.Inspect(file, astInspectNode(&results))
				}
			}
		}
	}

	return results, nil
}

func astInspectNode[S goast.Node](results *[]S) func(n goast.Node) bool {
	return func(n goast.Node) bool {
		switch impl := n.(type) {
		case S:
			*results = append(*results, impl)

			return false
		}

		return true
	}
}
