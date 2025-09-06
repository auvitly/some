package golang

import (
	"fmt"
	goast "go/ast"
	"go/token"
)

type inspectDeclGen struct{}

func (*inspectDeclGen) Toks(v any) ([]token.Token, error) {
	declList, err := any2node[*goast.GenDecl](v)
	if err != nil {
		return nil, err
	}

	var tokens []token.Token

	for _, declItem := range declList {
		tokens = append(tokens, declItem.Tok)
	}

	return tokens, nil
}

func (*inspectDeclGen) Tok(tok string, v any) ([]*goast.GenDecl, error) {
	if len(tok) == 0 {
		return nil, fmt.Errorf("token must be not empty")
	}

	declList, err := any2node[*goast.GenDecl](v)
	if err != nil {
		return nil, err
	}

	var decls []*goast.GenDecl

	for _, declItem := range declList {
		if declItem.Tok.String() == tok {
			decls = append(decls, declItem)
		}
	}

	return decls, nil
}
