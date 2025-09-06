package psql

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v6"
)

type inspect struct{}

func (funcs) Inspect() any { return inspect{} }

func (inspect) Select(v any) ([]*pg_query.SelectStmt, error) {
	var raws []*pg_query.RawStmt

	switch value := v.(type) {
	case *pg_query.ParseResult:
		raws = value.GetStmts()
	case *pg_query.RawStmt:
		raws = append(raws, value)
	default:
		return nil, fmt.Errorf("unsupported value for inspect")
	}

	var selectStmts []*pg_query.SelectStmt

	for _, raw := range raws {
		stmt := raw.GetStmt()
		if stmt == nil {
			continue
		}

		selectStmt := stmt.GetSelectStmt()
		if selectStmt == nil {
			continue
		}

		selectStmts = append(selectStmts, selectStmt)
	}

	return selectStmts, nil
}
