/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"io"
	"maps"
	"os"
	"os/exec"
	"reflect"
	"text/template"

	"github.com/auvitly/gopher/cmd"
	lib "github.com/auvitly/gopher/internal/template/libs"

	"gopkg.in/yaml.v3"
)

func main() {
	file, err := os.Open("templates/main.yaml")
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(2)
	}

	raw, err := io.ReadAll(file)
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(2)
	}

	var m = make(map[string]any)

	err = yaml.Unmarshal(raw, &m)
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(2)
	}

	fnMap := template.FuncMap{
		"typeof": func(v any) string { return reflect.TypeOf(v).Kind().String() },
		"shell": func(command string, args ...any) (any, error) {
			var _args []string

			for _, arg := range args {
				_args = append(_args, fmt.Sprintf("%s", arg))
			}

			res, err := exec.Command(command, _args...).Output()
			if err != nil {
				return nil, err
			}

			return string(res), nil
		},
		"error": func(args ...any) (any, error) {
			if len(args) == 0 {
				return nil, nil
			}

			if _, ok := args[0].(string); ok {
				return nil, fmt.Errorf(args[0].(string), args[1:]...)
			}

			return nil, fmt.Errorf("%v", args...)
		},
	}

	maps.Copy(fnMap, lib.Standard)

	tmpl, err := template.New("main").
		Funcs(fnMap).
		Parse(m["template"].(string))
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(2)
	}

	err = tmpl.Execute(os.Stdout, map[string]any{"Test": "test"})
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(2)
	}

	cmd.Execute()
}
