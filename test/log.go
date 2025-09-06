package test

import (
	sql "database/sql"

	uuid "github.com/google/uuid"
)

type Record struct {
	Field1 string `log:"-"`
	Field2 string `log:"*"`
	Field3 sql.NullBool
	Field4 Field4
	Field5 struct {
		SubField5Field string
	}
	Field6 uuid.NullUUID
}

type Field4 struct {
}
