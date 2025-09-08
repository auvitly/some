package engine

import (
	"errors"
)

func Error(format string, cond bool) (any, error) {
	if len(format) != 0 && cond {
		return nil, errors.New(format)
	}

	return "", nil
}
