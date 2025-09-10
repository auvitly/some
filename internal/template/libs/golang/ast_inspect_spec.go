package golang

import (
	goast "go/ast"
	"regexp"
)

type inspectSpec struct{}

// Types.

func (*inspectSpec) Types(v any) ([]*goast.TypeSpec, error) {
	return any2node[*goast.TypeSpec](v)
}

type inspectSpecType struct{}

// Types.

func (*inspectSpec) Type() any { return (*inspectSpecType)(nil) }

func (*inspectSpecType) Name(pattern string, v any) ([]*goast.TypeSpec, error) {
	specs, err := any2node[*goast.TypeSpec](v)
	if err != nil {
		return nil, err
	}

	var filtered []*goast.TypeSpec

	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	for _, spec := range specs {
		if spec.Name == nil {
			continue
		}

		if reg.Match([]byte(spec.Name.Name)) {
			filtered = append(filtered, spec)
		}
	}

	return filtered, nil
}

func (*inspectSpecType) StructType(v any) ([]*goast.TypeSpec, error) {
	return filterSpecTypeByType[*goast.StructType](v)
}

func (*inspectSpecType) InterfaceType(v any) ([]*goast.TypeSpec, error) {
	return filterSpecTypeByType[*goast.InterfaceType](v)
}

func (*inspectSpecType) Names(v any) ([]*goast.Ident, error) {
	specs, err := any2node[*goast.TypeSpec](v)
	if err != nil {
		return nil, err
	}

	var names []*goast.Ident

	for _, spec := range specs {
		names = append(names, spec.Name)
	}

	return names, nil
}

func (*inspectSpecType) Types(v any) ([]goast.Expr, error) {
	specs, err := any2node[*goast.TypeSpec](v)
	if err != nil {
		return nil, err
	}

	var names []goast.Expr

	for _, spec := range specs {
		if spec == nil {
			names = append(names, nil)
		} else {
			names = append(names, spec.Type)
		}
	}

	return names, nil
}

func filterSpecTypeByType[T goast.Expr](v any) ([]*goast.TypeSpec, error) {
	specs, err := any2node[*goast.TypeSpec](v)
	if err != nil {
		return nil, err
	}

	var list []*goast.TypeSpec

	for _, spec := range specs {
		if _, ok := spec.Type.(T); ok {
			list = append(list, spec)
		}
	}

	return list, nil
}

// Imports.

func (*inspectSpec) Imports(pattern any) ([]*goast.ImportSpec, error) {
	return any2node[*goast.ImportSpec](pattern)
}

type inspectSpecImport struct{}

func (*inspectSpec) Import() any { return (*inspectSpecImport)(nil) }

func (*inspectSpecImport) Name(pattern string, v any) ([]*goast.ImportSpec, error) {
	specs, err := any2node[*goast.ImportSpec](v)
	if err != nil {
		return nil, err
	}

	var filtered []*goast.ImportSpec

	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	for _, spec := range specs {
		if spec.Name == nil {
			continue
		}

		if reg.Match([]byte(spec.Name.Name)) {
			filtered = append(filtered, spec)
		}
	}

	return filtered, nil
}

func (*inspectSpecImport) Names(v any) ([]*goast.Ident, error) {
	specs, err := any2node[*goast.ImportSpec](v)
	if err != nil {
		return nil, err
	}

	var names []*goast.Ident

	for _, spec := range specs {
		names = append(names, spec.Name)
	}

	return names, nil
}

func (*inspectSpecImport) Path(pattern string, v any) ([]*goast.ImportSpec, error) {
	specs, err := any2node[*goast.ImportSpec](v)
	if err != nil {
		return nil, err
	}

	var filtered []*goast.ImportSpec

	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	for _, spec := range specs {
		if spec.Path == nil {
			continue
		}

		if reg.Match([]byte(spec.Path.Value)) {
			filtered = append(filtered, spec)
		}
	}

	return filtered, nil
}

func (*inspectSpecImport) Paths(v any) ([]*goast.BasicLit, error) {
	specs, err := any2node[*goast.ImportSpec](v)
	if err != nil {
		return nil, err
	}

	var names []*goast.BasicLit

	for _, spec := range specs {
		names = append(names, spec.Path)
	}

	return names, nil
}

func (*inspectSpec) Values(v any) ([]*goast.ValueSpec, error) {
	return any2node[*goast.ValueSpec](v)
}
