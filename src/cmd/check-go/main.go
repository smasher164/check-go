package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func main() {
	flag.Parse()
	fset := token.NewFileSet()
	file := flag.Arg(0)
	_, err := parser.ParseFile(fset, file, nil, 0)
	check(err)
}
