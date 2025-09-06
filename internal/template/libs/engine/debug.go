package engine

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/auvitly/gopher/internal/utils"
)

func Debug(v ...any) (any, error) {
	if len(v) == 0 {
		return nil, fmt.Errorf("nothing to debug")
	}

	rv := reflect.ValueOf(v[len(v)-1])

	var buf strings.Builder

	buf.WriteString("--- DEBUG\n")

	switch {
	case len(v) == 1 && v[0] != nil:
	case len(v) == 1 && v[0] == nil:
		buf.WriteString(fmt.Sprintf(">> Value: %v\n", v))
		buf.WriteString("---\n")

		return buf.String(), nil
	default:
		buf.WriteString("---\n")

		return buf.String(), nil
	}

	buf.WriteString("> Define\n")
	buf.WriteString(fmt.Sprintf(">> Type: %s\n", rv.Type().String()))

	if rv.Kind() == reflect.Pointer {
		buf.WriteString(fmt.Sprintf(">> Value: %v\n", utils.Dereference(rv).Interface()))
	} else {
		buf.WriteString(fmt.Sprintf(">> Value: %v\n", v))
	}

	switch utils.Dereference(rv).Kind() {
	case reflect.Struct:
		var str = utils.Dereference(rv)

		if str.NumField() != 0 {
			buf.WriteString("> Fields\n")
		}

		for i := range str.NumField() {
			rt := str.Type().Field(i)
			rv := str.Field(i)

			if !rt.IsExported() {
				continue
			}

			buf.WriteString(fmt.Sprintf(">> Field: %s | Type: %s | Value = %v\n", rt.Name, rt.Type.String(), rv.Interface()))
		}

		if rv.Type().NumMethod() != 0 {
			buf.WriteString("> Methods\n")
		}

		if rv.Kind() == reflect.Ptr {
			for i := range rv.Type().NumMethod() {
				method := rv.Type().Method(i)

				buf.WriteString(fmt.Sprintf(">> %s%s \n", method.Name, strings.TrimPrefix(method.Type.String(), "func")))
			}
		}
	}

	buf.WriteString("---\n")

	return buf.String(), nil
}
