package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "wire/integrate/main.go", nil, parser.AllErrors)
	ast.Print(fset, f)
}
