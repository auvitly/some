package test

import "encoding/json"

type Root struct {
	FieldValue   StructField  `json:"field_value"`
	FieldPointer *StructField `json:"field_pointer"`
	Marshaler    json.Marshaler
	FieldEmbed   struct {
		Field1 string
	}
}

type StructField struct {
}

func A() {

}

var _ = func() int {
	var a = 5

	return a
}()
