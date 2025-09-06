package os

import (
	"io"
	"os"
)

var Lib = func() any { return &lib{} }

type lib struct{}

func (*lib) PWD() (string, error) {
	return os.Getwd()
}

func (*lib) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func (*lib) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (*lib) Copy(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if srcInfo.IsDir() {
		return copyDir(src, dst)
	}

	return copyFile(src, dst)
}

func (*lib) Read(name string) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*lib) Write(name string, value any) (any, error) {
	return write(name, os.O_APPEND|os.O_WRONLY, value)
}

func (*lib) Overwrite(name string, value any) (any, error) {
	return write(name, os.O_TRUNC|os.O_WRONLY, value)
}

func (*lib) Mkdir(name string) error {
	return os.Mkdir(name, os.ModeDir)
}

func (*lib) Remove(name string) error {
	return os.RemoveAll(name)
}

func (*lib) Stdout() io.Writer {
	return os.Stdout
}

func (*lib) Stdin() io.Writer {
	return os.Stdin
}

func (*lib) Stderr() io.Writer {
	return os.Stderr
}

// Shortcuts.

func (l *lib) OW(name string, value any) (any, error) {
	return l.Overwrite(name, value)
}
