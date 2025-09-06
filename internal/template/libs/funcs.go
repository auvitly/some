package libs

import (
	"reflect"
	"text/template"

	"github.com/auvitly/gopher/internal/template/libs/engine"
	"github.com/auvitly/gopher/internal/template/libs/env"
	"github.com/auvitly/gopher/internal/template/libs/golang"
	"github.com/auvitly/gopher/internal/template/libs/io"
	"github.com/auvitly/gopher/internal/template/libs/json"
	"github.com/auvitly/gopher/internal/template/libs/maps"
	"github.com/auvitly/gopher/internal/template/libs/math"
	"github.com/auvitly/gopher/internal/template/libs/obj"
	"github.com/auvitly/gopher/internal/template/libs/os"
	"github.com/auvitly/gopher/internal/template/libs/psql"
	"github.com/auvitly/gopher/internal/template/libs/variable"
)

var Standard = template.FuncMap{
	// 1 Format:
	// 1.1 Vold/Debug
	"doc":     engine.Doc,
	"log":     engine.Log,
	"void":    engine.Void,
	"debug":   engine.Debug,
	"dump":    engine.Dump,
	"typeof":  reflect.TypeOf,
	"valueof": reflect.ValueOf,
	// 1.2 Symbols
	"ln":    engine.Line,
	"tab":   engine.Tab,
	"dot":   engine.Dot,
	"space": engine.Space,

	// 2 Types
	// 2.1 Base
	"null":   engine.Null,
	"slice":  engine.Slice,
	"string": engine.String,
	"int":    engine.Int,
	"float":  engine.Float,
	"bytes":  engine.Bytes,
	"bool":   engine.Bool,
	"map":    engine.Map,
	// 2.2 Funcs for types
	"unique":   engine.Unique,
	"contains": engine.Contains,
	"append":   engine.Append,
	"index":    engine.Index,
	"coalesce": engine.Coalesce,
	"set":      engine.Set,
	"get":      engine.Get,

	// 3 Extensions
	// 3.1 General
	"maps": maps.Lib,
	"obj":  obj.Lib,
	"os":   os.Lib,
	"io":   io.Lib,
	"json": json.Lib,
	"go":   golang.Lib,
	"env":  env.Lib,
	"math": math.Lib,
	"psql": psql.Lib,
	"var":  variable.Lib,
}
