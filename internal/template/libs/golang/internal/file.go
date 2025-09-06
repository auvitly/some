package internal

import (
	"bytes"
	goast "go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

type File struct {
	GoMod  *modfile.File
	Fset   *token.FileSet
	Syntax *goast.File
}

func (f *File) String() (string, error) {
	var buf = bytes.NewBuffer(nil)

	err := format.Node(buf, f.Fset, f.Syntax)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ScanFile(path string) (*File, error) {
	asbPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	goModPath, err := FindGoMod(asbPath)
	if err != nil {
		return nil, err
	}

	modData, err := os.ReadFile(goModPath)
	if err != nil {
		return nil, err
	}

	var file = new(File)

	file.GoMod, err = modfile.Parse(goModPath, modData, nil)
	if err != nil {
		return nil, err
	}

	file.Fset = token.NewFileSet()

	file.Syntax, err = parser.ParseFile(file.Fset, asbPath, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return file, nil
}
