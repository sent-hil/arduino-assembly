package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

var (
	emptyFile       = "examples/empty.go"
	forFile         = "examples/for.go"
	funcCallFile    = "examples/func_call.go"
	forFuncCallFile = "examples/for_func_call.go"
	funcDefFile     = "examples/func_def.go"
	blinkFile       = "examples/blink.go"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, emptyFile, nil, 0)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, f)
}
