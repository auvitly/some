package golang

import (
	"fmt"
	goast "go/ast"
	"regexp"
)

type inspectDeclFunc struct{}

// Greps.

func (*inspectDeclFunc) Names(v any) ([]*goast.Ident, error) {
	var names []*goast.Ident

	funcDecls, err := any2node[*goast.FuncDecl](v)
	if err != nil {
		return nil, err
	}

	for _, funcDecl := range funcDecls {
		names = append(names, funcDecl.Name)
	}

	return names, nil
}

func (*inspectDeclFunc) CommentGroups(v any) ([]*goast.CommentGroup, error) {
	var comments []*goast.CommentGroup

	funcDecls, err := any2node[*goast.FuncDecl](v)
	if err != nil {
		return nil, err
	}

	for _, funcDecl := range funcDecls {
		comments = append(comments, funcDecl.Doc)
	}

	return comments, nil
}

func (*inspectDeclFunc) Recvs(v any) ([]*goast.FieldList, error) {
	var recvs []*goast.FieldList

	funcDecls, err := any2node[*goast.FuncDecl](v)
	if err != nil {
		return nil, err
	}

	for _, funcDecl := range funcDecls {
		recvs = append(recvs, funcDecl.Recv)
	}

	return recvs, nil
}

func (*inspectDeclFunc) Recv(pattern string, v any) ([]*goast.FuncDecl, error) {
	var decls []*goast.FuncDecl

	funcDecls, err := any2node[*goast.FuncDecl](v)
	if err != nil {
		return nil, err
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	for _, funcDecl := range funcDecls {
		if funcDecl.Recv == nil {
			continue
		}

		if len(funcDecl.Recv.List) != 1 {
			continue
		}

		if len(funcDecl.Recv.List[0].Names) != 0 {
			continue
		}

		if !re.MatchString(funcDecl.Recv.List[0].Names[0].Name) {
			continue
		}

		decls = append(decls, funcDecl)
	}

	return decls, nil
}

func (*inspectDeclFunc) Types(v any) ([]*goast.FuncType, error) {
	var types []*goast.FuncType

	funcDecls, err := any2node[*goast.FuncDecl](v)
	if err != nil {
		return nil, err
	}

	for _, funcDecl := range funcDecls {
		types = append(types, funcDecl.Type)
	}

	return types, nil
}

func (*inspectDeclFunc) Bodys(v any) ([]*goast.BlockStmt, error) {
	var bodys []*goast.BlockStmt

	funcDecls, err := any2node[*goast.FuncDecl](v)
	if err != nil {
		return nil, err
	}

	for _, funcDecl := range funcDecls {
		bodys = append(bodys, funcDecl.Body)
	}

	return bodys, nil
}

// Filters.

func (*inspectDeclFunc) Exported(v any) ([]*goast.FuncDecl, error) {
	return any2FuncDecls(func(fd *goast.FuncDecl) bool {
		return fd.Name.IsExported()
	}, v)
}

func (*inspectDeclFunc) Unexported(v any) ([]*goast.FuncDecl, error) {
	return any2FuncDecls(func(fd *goast.FuncDecl) bool {
		return !fd.Name.IsExported()
	}, v)
}

func (*inspectDeclFunc) Name(pattern string, v any) ([]*goast.FuncDecl, error) {
	if len(pattern) == 0 {
		return nil, fmt.Errorf("pattern must be not empty")
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return any2FuncDecls(func(fd *goast.FuncDecl) bool {
		return re.MatchString(fd.Name.String())
	}, v)
}

func (*inspectDeclFunc) ReceiverType(pattern string, v any) ([]*goast.FuncDecl, error) {
	if len(pattern) == 0 {
		return nil, fmt.Errorf("pattern must be not empty")
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	var fn = func(fd *goast.FuncDecl) bool {
		if fd.Recv == nil {
			return false
		}

		for _, item := range fd.Recv.List {
			if item.Type == nil {
				continue
			}

			obj, ok := item.Type.(*goast.Ident)
			if ok {
				return re.MatchString(obj.String())
			}

			star, ok := item.Type.(*goast.StarExpr)
			if !ok {
				return false
			}

			obj, ok = star.X.(*goast.Ident)
			if ok {
				return re.MatchString(obj.String())
			}
		}

		return false
	}

	return any2FuncDecls(fn, v)
}

func (*inspectDeclFunc) ReceiverPointer(v any) ([]*goast.FuncDecl, error) {
	var fn = func(fd *goast.FuncDecl) bool {
		if fd.Recv == nil {
			return false
		}

		for _, item := range fd.Recv.List {
			if item.Type == nil {
				return false
			}

			_, ok := item.Type.(*goast.StarExpr)
			if ok {
				return true
			}
		}

		return false
	}

	return any2FuncDecls(fn, v)
}

func (*inspectDeclFunc) ReceiverValue(v any) ([]*goast.FuncDecl, error) {
	var fn = func(fd *goast.FuncDecl) bool {
		if fd.Recv == nil {
			return false
		}

		for _, item := range fd.Recv.List {
			if item.Type == nil {
				continue
			}

			_, ok := item.Type.(*goast.StarExpr)
			if !ok {
				return true
			}
		}

		return false
	}

	return any2FuncDecls(fn, v)
}

func any2FuncDecls(fn func(*goast.FuncDecl) bool, v any) ([]*goast.FuncDecl, error) {
	var decls []*goast.FuncDecl

	funcDecls, err := any2node[*goast.FuncDecl](v)
	if err != nil {
		return nil, err
	}
	for _, funcDecl := range funcDecls {
		if fn(funcDecl) {
			decls = append(decls, funcDecl)
		}
	}

	return decls, nil
}
