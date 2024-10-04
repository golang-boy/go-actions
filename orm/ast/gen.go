package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
)

func Gen(w io.Writer, file string) error {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	s := &Visitor{}
	ast.Walk(s, f)

	ff := s.Get()

	fmt.Print(ff)

	return nil

}
