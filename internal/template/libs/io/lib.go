package io

import "io"

var Lib = func() any { return &lib{} }

type lib struct{}

func (*lib) ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}
