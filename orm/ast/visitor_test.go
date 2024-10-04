package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVisitor(t *testing.T) {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", `package ast

import (
	"fmt"
	"go/ast"
	"reflect"
)

type Visitor struct {
}

func (p Visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return p
	}

	typ := reflect.TypeOf(node)

	val := reflect.ValueOf(node)

	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	fmt.Printf("val: %v, typ: %s", val.Interface(), typ.Name())

	return p

}
`, parser.ParseComments)
	require.NoError(t, err)

	v := &Visitor{}
	ast.Walk(v, f)
}
