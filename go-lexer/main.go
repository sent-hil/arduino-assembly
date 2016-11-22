package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	src := []byte("// something something\n// something something\npackage main\nfunc main() {\nfmt.Println()}\n")

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	// Traverse the input source
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
