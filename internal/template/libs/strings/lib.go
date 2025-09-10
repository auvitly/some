package strings

import (
	"strconv"
	"strings"

	"github.com/auvitly/gopher/internal/utils"
)

var Lib = func() any { return &lib{} }

type lib struct{}

func (*lib) Unquote(v ...any) (any, error) {
	return any2stringFunc(func(s string) (any, error) {
		return strconv.Unquote(s)
	}, v)
}

func (*lib) Quote(v ...any) (any, error) {
	return any2stringFunc(func(s string) (any, error) {
		return strconv.Quote(s), nil
	}, v)
}

func (*lib) Trim(str string, v ...any) (any, error) {
	return any2stringFunc(func(s string) (any, error) {
		return strings.Trim(s, str), nil
	}, v)
}

func (*lib) CutPrefix(str string, v ...any) (any, error) {
	return any2stringFunc(func(s string) (any, error) {
		s, _ = strings.CutPrefix(s, str)

		return s, nil
	}, v)
}

func any2stringFunc(fn func(string) (any, error), v []any) (any, error) {
	strs, err := utils.AnyTo[string](v...)
	if err != nil {
		return nil, err
	}

	var results = make([]any, len(strs))

	for i := range strs {
		results[i], err = fn(strs[i])
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}
