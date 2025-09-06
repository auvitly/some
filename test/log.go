package test

import (
	sql "database/sql"

	uuid "github.com/google/uuid"
)

type Record struct {
	Field1, Field2 string `log:"-"`
	Field3         string `log:"*"`
	Field4         sql.NullBool
	Field5         Field5
	Field6         struct {
		SubField5Field string
	}
	Field7 uuid.NullUUID
}

type Field5 struct {
}
