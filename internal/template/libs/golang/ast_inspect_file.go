package golang

import (
	"fmt"
	goast "go/ast"
	"regexp"

	"github.com/auvitly/gopher/internal/template/libs/golang/internal"
	"golang.org/x/tools/go/packages"
)

type inspectFile struct{}

type matchFileFunc func(file *goast.File) bool

var inspectFileImpl = new(inspectFile)

func (*inspectFile) PackageName(pattern string, v any) ([]*goast.File, error) {
	return any2Files(matchFilePackageName(pattern), v)
}

func (*inspectFile) Test(v any) ([]*goast.File, error) {
	return any2Files(matchFileTest(), v)
}

func any2Files(fn matchFileFunc, v any) (list []*goast.File, _ error) {
	switch value := v.(type) {
	case *internal.Project:
		return packages2Files(fn, value.Packages...)
	case []*packages.Package:
		return packages2Files(fn, value...)
	case *packages.Package:
		return packages2Files(fn, value)
	case []*goast.File:
		return files2Files(fn, value...)
	case *goast.File:
		return files2Files(fn, value)
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}
}

func packages2Files(fn matchFileFunc, list ...*packages.Package) ([]*goast.File, error) {
	var files []*goast.File

	for _, item := range list {
		results, err := files2Files(fn, item.Syntax...)
		if err != nil {
			return nil, err
		}

		files = append(files, results...)
	}

	return files, nil
}

func files2Files(fn matchFileFunc, list ...*goast.File) ([]*goast.File, error) {
	var files []*goast.File

	for _, item := range list {
		if fn(item) {
			files = append(files, item)
		}
	}

	if len(files) == 0 {
		return nil, nil
	}

	return files, nil
}

func matchFilePackageName(pattern string) matchFileFunc {
	return func(file *goast.File) bool {
		var matchRegExpName, _ = regexp.MatchString(pattern, file.Name.String())

		return matchRegExpName
	}
}

func matchFileTest() matchFileFunc {
	return func(file *goast.File) bool {
		var matchRegExpName, _ = regexp.MatchString("_test", file.Name.String())

		return matchRegExpName
	}
}
