package ast

import (
	"fmt"
	"go/ast"
)

type Visitor struct {
	file *FileVisitor
}

func (p *Visitor) Get() *File {
	return &File{
		Package: p.file.Package,
	}
}

func (p Visitor) Visit(node ast.Node) (w ast.Visitor) {

	fn, ok := node.(*ast.File)
	fmt.Println(ok)
	fmt.Println(fn.Name)

	if !ok {
		return p
	}

	p.file = &FileVisitor{
		Package: fn.Name.String(),
	}

	return p.file
}

type File struct {
	Package string
}

type FileVisitor struct {
	Package string
}

func (f *FileVisitor) Visit(node ast.Node) (w ast.Visitor) {
	return f
}
