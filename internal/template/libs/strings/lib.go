package strings

import (
	"strconv"

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
