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

func (*inspectSpec) Imports(pattern any) ([]*goast.ImportSpec, error) {
	return any2node[*goast.ImportSpec](pattern)
}

func (*inspectSpec) Values(v any) ([]*goast.ValueSpec, error) {
	return any2node[*goast.ValueSpec](v)
}
