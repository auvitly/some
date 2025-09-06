package golang

import (
	goast "go/ast"
	"go/token"
)

type inspectPosition struct{}

func (*inspectPosition) Filename(fset *token.FileSet, node goast.Node) (string, error) {
	return fset.Position(node.Pos()).Filename, nil
}

func (*inspectPosition) Column(fset *token.FileSet, node goast.Node) (int, error) {
	return fset.Position(node.Pos()).Column, nil
}

func (*inspectPosition) Line(fset *token.FileSet, node goast.Node) (int, error) {
	return fset.Position(node.Pos()).Line, nil
}

func (*inspectPosition) LineEnd(fset *token.FileSet, node goast.Node) (int, error) {
	return fset.Position(node.End()).Line, nil
}

func (*inspectPosition) Node(fset *token.FileSet, line int, file *goast.File) (goast.Node, error) {
	var node goast.Node

	goast.Inspect(file, func(n goast.Node) bool {
		if n == nil {
			return false
		}

		if _, ok := n.(*goast.File); ok {
			return true
		}

		var posIter = fset.Position(n.Pos())

		if posIter.Line != line {
			return true
		}

		if node == nil {
			node = n
		}

		return true
	})

	return node, nil
}
