package test

import (
	"database/sql"

	"github.com/google/uuid"
)

type (
	String  string
	Int     int
	Pointer *sql.NullBool
	Import  sql.NullBool
	Alias   = Field5
	Adapter Field5
	Chan    chan string
	Map     map[string]sql.NullBool
	Slice   []any
	Array   [3]any
)

type Record struct {
	Field1, Field2 string `log:"-" mt:"test"`
	Field3         string `log:"*"`
	Field4         sql.NullBool
	Field5         Field5
	Field6         struct {
		SubField5Field string
	}
	Field7 uuid.NullUUID
	Field8 uuid.NullUUID
}

type Field5 struct {
}
