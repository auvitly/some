package internal

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"
)

type Project struct {
	GoMod    *modfile.File
	Packages []*packages.Package
}

func (p *Project) String() (string, error) {
	var buf = bytes.NewBuffer(nil)

	for _, pkg := range p.Packages {
		for _, file := range pkg.Syntax {
			err := format.Node(buf, pkg.Fset, file)
			if err != nil {
				return "", err
			}
		}
	}

	return buf.String(), nil
}

func ScanProject(path string) (*Project, error) {
	gomodPath, err := FindGoMod(path)
	if err != nil {
		return nil, err
	}

	modData, err := os.ReadFile(gomodPath)
	if err != nil {
		return nil, err
	}

	var project = new(Project)

	project.GoMod, err = modfile.Parse(gomodPath, modData, nil)
	if err != nil {
		return nil, err
	}

	project.Packages, err = ScanPackages(fmt.Sprintf("%s/...", filepath.Dir(gomodPath)))
	if err != nil {
		return nil, err
	}

	return project, nil
}

func FindGoMod(startDir string) (string, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}

	dir = filepath.Dir(dir)

	for {
		goModPath := filepath.Join(dir, "go.mod")

		if _, err := os.Stat(goModPath); err == nil {
			return goModPath, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return "", fmt.Errorf("go.mod not found")
		}

		dir = parentDir
	}
}
