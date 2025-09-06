package engine

import (
	"bytes"
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// Dump расширенный вывод для дебага с поддержкой циклов и типов
func Dump(obj any) any {
	return dumpWithConfig(obj)
}

type Visited map[VisitedKey]int

type VisitedKey struct {
	Addr uintptr
	Type string
}

func dumpWithConfig(obj any) any {
	var buf = bytes.NewBuffer(nil)

	fmt.Fprintf(buf, "╔═ DUMP ═\n")
	fmt.Fprintf(buf, "║ Type:\t%T\n", obj)
	fmt.Fprintf(buf, "╠═ Value:\n")

	var prefix = "║ "

	fmt.Fprintf(buf, "%s", prefix)

	var visited = make(Visited)

	dumpValue(buf, reflect.ValueOf(obj), "║ ", visited)

	fmt.Fprintf(buf, "\n")

	fmt.Fprintf(buf, "╚════════\n")

	return buf.String()
}

func dumpValue(buf *bytes.Buffer, v reflect.Value, prefix string, visited Visited) {
	if !v.IsValid() {
		fmt.Fprintf(buf, "<invalid>")

		return
	}

	if v.CanAddr() {
		var ptr = v.UnsafeAddr()

		if visited[VisitedKey{Addr: ptr, Type: v.Type().String()}] > 0 {
			fmt.Fprintf(buf, "&%p", unsafe.Pointer(v.UnsafeAddr()))

			return
		}

		visited[VisitedKey{Addr: ptr, Type: v.Type().String()}]++
	}

	switch v.Kind() {
	case reflect.Struct:
		dumpStruct(buf, v, prefix, visited)
	case reflect.Ptr:
		dumpPointer(buf, v, prefix, visited)
	case reflect.Interface:
		dumpInterface(buf, v, prefix, visited)
	case reflect.Slice:
		dumpSlice(buf, v, prefix, visited)
	case reflect.Array:
		dumpArray(buf, v, prefix, visited)
	case reflect.Map:
		dumpMap(buf, v, prefix, visited)
	case reflect.String:
		var str = v.String()

		if utf8.ValidString(str) {
			fmt.Fprintf(buf, "%q", str)
		} else {
			fmt.Fprintf(buf, "%v (binary data)", []byte(str))
		}
	default:
		if v.CanInterface() {
			fmt.Fprintf(buf, "%v", v.Interface())
		}
	}
}

func dumpStruct(buf *bytes.Buffer, v reflect.Value, prefix string, visited Visited) {
	var (
		t        = v.Type()
		typename = value2typename(v)
	)

	fmt.Fprintf(buf, "%s {\n", typename)

	for i := range v.NumField() {
		var (
			fieldType  = t.Field(i)
			fieldValue = v.Field(i)
		)

		if fieldType.IsExported() {
			var fieldTypeInfo = fmt.Sprintf("<%s>", getFieldTypeInfo(fieldValue))

			fmt.Fprintf(buf, "%s\t%s %s:", prefix, fieldType.Name, fieldTypeInfo)

			dumpValue(buf, fieldValue, prefix+"\t", visited)

			fmt.Fprintf(buf, "\n")
		}
	}

	// if v.NumField() == 0 {
	// 	fmt.Fprintf(buf, "\n")
	// }

	fmt.Fprintf(buf, "%s}", prefix)
}

func dumpPointer(buf *bytes.Buffer, v reflect.Value, prefix string, visited Visited) {
	var typename = value2typename(v)

	switch {
	case v.IsNil():
		fmt.Fprintf(buf, "%s(nil)", typename)
	default:
		if v.CanInterface() {
			var addrInfo = fmt.Sprintf("@%p -> ", v.Interface())

			if v.Elem().Kind() == reflect.Pointer && v.Elem().IsNil() {
				fmt.Fprintf(buf, "%s %s", typename, addrInfo)
				fmt.Fprintf(buf, "(nil)")
			} else {
				if v.Elem().CanAddr() {
					var ptr = v.Elem().UnsafeAddr()

					if visited[VisitedKey{Addr: ptr, Type: v.Elem().Type().String()}] > 0 {
						fmt.Fprintf(buf, "&%p", unsafe.Pointer(v.Elem().UnsafeAddr()))

						return
					}

					visited[VisitedKey{Addr: ptr, Type: v.Type().String()}]++
				}

				fmt.Fprintf(buf, "%s %s", typename, addrInfo)

				dumpValue(buf, v.Elem(), prefix, visited)
			}
		} else {
			fmt.Fprintf(buf, "<inaccessible>")
		}
	}
}

func dumpInterface(buf *bytes.Buffer, v reflect.Value, prefix string, visited Visited) {
	if v.IsNil() {
		fmt.Fprintf(buf, "<nil>")
	} else {
		dumpValue(buf, v.Elem(), prefix, visited)
	}
}

func dumpMap(buf *bytes.Buffer, v reflect.Value, prefix string, visited Visited) {
	if v.IsNil() {
		fmt.Fprintf(buf, "<nil>\n")
	} else {
		fmt.Fprintf(buf, "{\n")

		for _, key := range v.MapKeys() {
			fmt.Fprintf(buf, "%s\t[", prefix)

			dumpValue(buf, key, "\t", visited)

			fmt.Fprintf(buf, "]: ")

			dumpValue(buf, v.MapIndex(key), prefix+"\t", visited)

			fmt.Fprintf(buf, "\n")
		}

		// if v.Len() == 0 {
		// 	fmt.Fprintf(buf, "\n")
		// }

		fmt.Fprintf(buf, "%s}", prefix)
	}
}

func dumpSlice(buf *bytes.Buffer, v reflect.Value, prefix string, visited Visited) {
	if v.IsNil() {
		fmt.Fprintf(buf, "<nil>")
	} else {
		if v.Len() == 0 {
			fmt.Fprintf(buf, "[]")

			return
		}

		fmt.Fprintf(buf, "[\n")

		for i := 0; i < v.Len(); i++ {
			fmt.Fprintf(buf, "%s\t[%d]:", prefix, i)

			dumpValue(buf, v.Index(i), prefix+"\t", visited)

			fmt.Fprintf(buf, "\n")
		}

		fmt.Fprintf(buf, "%s]", prefix)
	}
}

func dumpArray(buf *bytes.Buffer, v reflect.Value, prefix string, visited Visited) {
	if v.Len() == 0 {
		fmt.Fprintf(buf, "[]")

		return
	}

	fmt.Fprintf(buf, "[\n")

	for i := 0; i < v.Len(); i++ {
		fmt.Fprintf(buf, "%s\t[%d]:", prefix, i)

		dumpValue(buf, v.Index(i), prefix+"\t", visited)

		fmt.Fprintf(buf, "\n")
	}

	fmt.Fprintf(buf, "%s]", prefix)
}

func value2typename(v reflect.Value) string {
	if !v.IsValid() {
		return "<invalid>"
	}

	kind := v.Kind()

	switch kind {
	case reflect.Ptr:
		if v.IsNil() {
			return fmt.Sprintf("<%s>", v.Type())
		}

		return fmt.Sprintf("<*%s>", v.Elem().Type())
	case reflect.Interface:
		if v.IsNil() {
			return fmt.Sprintf("<%s>", v.Type())
		}

		return fmt.Sprintf("<%s:%s>", v.Type(), v.Elem().Type())
	default:
		return fmt.Sprintf("<%s>", v.Type())
	}
}

func getFieldTypeInfo(v reflect.Value) string {
	if !v.IsValid() {
		return "invalid"
	}

	if v.Kind() == reflect.Interface && !v.IsNil() {
		return fmt.Sprintf("%s~%s", v.Type(), v.Elem().Type())
	}

	return v.Type().String()
}
