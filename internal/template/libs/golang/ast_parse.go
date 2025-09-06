package golang

import (
	"fmt"
	"os"
	"reflect"

	"github.com/auvitly/gopher/internal/template/libs/golang/internal"
	"golang.org/x/tools/go/packages"
)

type parse struct{}

func (*parse) Project(paths ...string) (any, error) {
	switch len(paths) {
	case 0:
		return internal.ScanProject("./...")
	case 1:
		return internal.ScanProject(paths[0])
	default:
		var projects []*internal.Project

		for _, path := range paths {
			project, err := internal.ScanProject(path)
			if err != nil {
				return nil, err
			}

			projects = append(projects, project)
		}

		return projects, nil
	}
}

func (*parse) Packages(path ...string) ([]*packages.Package, error) {
	return internal.ScanPackages(path...)
}

func (*parse) File(path string) (any, error) {
	switch len(path) {
	case 0:
		if len(os.Getenv("GOFILE")) == 0 {
			return nil, fmt.Errorf("not found path for parsing file")
		}

		return internal.ScanFile(os.Getenv("GOFILE"))
	default:
		return internal.ScanFile(path)
	}
}

func (*parse) Files(values ...any) (any, error) {
	var files []*internal.File

	switch len(values) {
	case 0:
		if len(os.Getenv("GOFILE")) == 0 {
			return nil, fmt.Errorf("not found path for parsing file")
		}

		file, err := internal.ScanFile(os.Getenv("GOFILE"))
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	default:
		for _, value := range values {
			var rv = reflect.ValueOf(value)

			var list []any

			switch rv.Kind() {
			case reflect.Slice:
				for i := range rv.Len() {
					list = append(list, rv.Index(i).Interface())
				}
			default:
				list = append(list, rv.Interface())
			}

			for _, item := range list {
				file, err := internal.ScanFile(fmt.Sprintf("%v", item))
				if err != nil {
					return nil, err
				}

				files = append(files, file)
			}
		}
	}

	return files, nil
}
