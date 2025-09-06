package os

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func write(name string, flag int, value any) error {
	dir := filepath.Dir(name)

	// Создаем все необходимые директории (с правами 0755)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(name, flag, 777)
	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(name)
		if err != nil {
			return err
		}
	}

	defer file.Close()

	var data []byte

	switch v := value.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	case interface{ String() string }:
		data = []byte(v.String())
	case interface{ String() (string, error) }:
		str, err := v.String()
		if err != nil {
			return err
		}

		data = []byte(str)
	default:
		data = fmt.Appendf(nil, "%v", v)
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
