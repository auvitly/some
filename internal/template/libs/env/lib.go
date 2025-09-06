package env

import (
	"os"
	"strings"
)

var Lib = func() any { return &lib{} }

type lib struct{}

func (*lib) Set(key string, value string) error {
	return os.Setenv(key, value)
}

func (*lib) List() map[string]any {
	var m = make(map[string]any)

	for _, value := range os.Environ() {
		parts := strings.Split(value, "=")

		m[parts[0]] = os.Getenv(parts[0])
	}

	return m
}

func (*lib) Get(key string) any {
	return os.Getenv(key)
}
