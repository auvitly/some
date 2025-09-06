package golang

import (
	"fmt"
	"regexp"

	"github.com/auvitly/gopher/internal/template/libs/golang/internal"
	"golang.org/x/tools/go/packages"
)

type inspectPackage struct{}

type matchPackageFunc func(pkg *packages.Package) bool

func (*inspectPackage) ID(pattern string, v any) ([]*packages.Package, error) {
	return any2Packages(matchPackageID(pattern), v)
}

func (*inspectPackage) Name(pattern string, v any) ([]*packages.Package, error) {
	return any2Packages(matchPackageName(pattern), v)
}

func (*inspectPackage) Path(pattern string, v any) ([]*packages.Package, error) {
	return any2Packages(matchPackagePath(pattern), v)
}

func (*inspectPackage) Dir(pattern string, v any) ([]*packages.Package, error) {
	return any2Packages(matchPackageDir(pattern), v)
}

func (*inspectPackage) Test(v any) ([]*packages.Package, error) {
	return any2Packages(matchPackageName("_test"), v)
}

func (*inspectPackage) WithoutTest(v any) ([]*packages.Package, error) {
	return any2Packages(matchWithoutTest("_test"), v)
}

func any2Packages(fn matchPackageFunc, v any) ([]*packages.Package, error) {
	switch value := v.(type) {
	case *internal.Project:
		return packages2packages(fn, value.Packages...)
	case []*packages.Package:
		return packages2packages(fn, value...)
	case *packages.Package:
		return packages2packages(fn, value)
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}
}

func packages2packages(fn matchPackageFunc, list ...*packages.Package) ([]*packages.Package, error) {
	var pkgs []*packages.Package

	for _, item := range list {
		if fn(item) {
			pkgs = append(pkgs, item)
		}
	}

	return pkgs, nil
}

func matchPackageID(pattern string) matchPackageFunc {
	return func(pkg *packages.Package) bool {
		var matchRegExpName, _ = regexp.MatchString(pattern, pkg.ID)

		return matchRegExpName
	}
}

func matchPackageName(pattern string) matchPackageFunc {
	return func(pkg *packages.Package) bool {
		var matchRegExpName, _ = regexp.MatchString(pattern, pkg.Name)

		return matchRegExpName
	}
}

func matchPackagePath(path string) matchPackageFunc {
	return func(pkg *packages.Package) bool {
		var matchRegExpName, _ = regexp.MatchString(path, pkg.PkgPath)

		return matchRegExpName
	}
}

func matchPackageDir(path string) matchPackageFunc {
	return func(pkg *packages.Package) bool {
		var matchRegExpName, _ = regexp.MatchString(path, pkg.Dir)

		return matchRegExpName
	}
}

func matchWithoutTest(path string) matchPackageFunc {
	return func(pkg *packages.Package) bool {
		var matchRegExpName, _ = regexp.MatchString(path, pkg.Dir)

		return !matchRegExpName
	}
}
