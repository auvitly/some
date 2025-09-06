package internal

import (
	"errors"
	"os"

	"golang.org/x/tools/go/packages"
)

func ScanPackages(patterns ...string) ([]*packages.Package, error) {
	if len(patterns) == 0 {
		patterns = append(patterns, "./...")
	}

	packages, err := packages.Load(
		&packages.Config{
			Mode:  packages.LoadFiles | packages.LoadImports | packages.LoadAllSyntax | packages.NeedForTest,
			Dir:   os.Getenv("PWD"),
			Tests: true,
		},
		patterns...,
	)
	if err != nil {
		return nil, err
	}

	var errs []error

	for _, pkg := range packages {
		for _, err := range pkg.Errors {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return nil, errors.Join(errs...)
	}

	return packages, nil
}
