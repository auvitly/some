package engine

import (
	"bytes"
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

	var buf = bytes.NewBuffer(nil)

	fmt.Fprintf(buf, "╔═ DEBUG ═\n")

	switch {
	case len(v) == 1 && v[0] != nil:
	case len(v) == 1 && v[0] == nil:
		fmt.Fprintf(buf, "║ Type <%v>\n", reflect.TypeOf(v).String())
		fmt.Fprintf(buf, "║ Value %v\n", v)
		fmt.Fprintf(buf, "╚═════════\n")

		return buf.String(), nil
	default:
		fmt.Fprintf(buf, "╚═════════\n")

		return buf.String(), nil
	}

	fmt.Fprintf(buf, "║ Type: <%v>\n", rv.Type().String())

	if rv.Kind() == reflect.Pointer {
		fmt.Fprintf(buf, "║ Value: %v\n", utils.Dereference(rv).Interface())
	} else {
		fmt.Fprintf(buf, "║ Value: %v\n", v)
	}

	switch utils.Dereference(rv).Kind() {
	case reflect.Struct:
		var str = utils.Dereference(rv)

		if str.NumField() != 0 {
			buf.WriteString("╠═ Fields\n")
		}

		for i := range str.NumField() {
			rt := str.Type().Field(i)
			rv := str.Field(i)

			if !rt.IsExported() {
				continue
			}

			fmt.Fprintf(buf, "║ %s <%s>:%v\n", rt.Name, rt.Type.String(), rv.Interface())
		}

		if rv.Type().NumMethod() != 0 {
			fmt.Fprintf(buf, "╠═ Methods\n")
		}

		if rv.Kind() == reflect.Ptr {
			for i := range rv.Type().NumMethod() {
				method := rv.Type().Method(i)

				fmt.Fprintf(buf, "║ %s%s \n", method.Name, strings.TrimPrefix(method.Type.String(), "func"))
			}
		}
	}

	fmt.Fprintf(buf, "╚═════════\n")

	return buf.String(), nil
}
