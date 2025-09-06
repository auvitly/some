package psql

import (
	"fmt"
	"io"
	"os"

	pg_query "github.com/pganalyze/pg_query_go/v6"
)

var Lib = func() any { return funcs{} }

type funcs struct{}

func (funcs) Parse(v any) (any, error) {
	var query string

	switch value := v.(type) {
	case string:
		query = value
	case []byte:
		query = string(query)
	default:
		return nil, fmt.Errorf("invalid argument type")
	}

	tree, err := pg_query.Parse(query)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

func (funcs) Read(name string) (any, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	tree, err := pg_query.Parse(string(data))
	if err != nil {
		return nil, err
	}

	return tree, nil
}
