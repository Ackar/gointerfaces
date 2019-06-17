package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func getFolderInterfaces(dir string) ([]string, error) {
	var fset token.FileSet
	res, err := parser.ParseDir(&fset, dir, nil, 0)
	if err != nil {
		return nil, err
	}

	var interfaces []string
	for pkg, pkgAst := range res {
		for _, a := range pkgAst.Files {
			for _, decl := range a.Decls {
				switch v := decl.(type) {
				case *ast.GenDecl:
					if v.Tok == token.TYPE {
						ts := v.Specs[0].(*ast.TypeSpec)

						if _, ok := ts.Type.(*ast.InterfaceType); ok {
							if !ts.Name.IsExported() {
								continue
							}
							interfaces = append(interfaces, pkg+"."+ts.Name.Name)
						}
					}
				}
			}
		}
	}

	return interfaces, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s dir1 dir2 ...\n", os.Args[0])
		os.Exit(1)
	}

	var interfaces []string
	for _, arg := range os.Args[1:] {
		intfcs, err := getFolderInterfaces(arg)
		if err != nil {
			fmt.Printf("Error while getting interfaces for %s: %s\n", arg, err)
			os.Exit(2)
		}
		interfaces = append(interfaces, intfcs...)
	}

	for _, i := range interfaces {
		fmt.Println(i)
	}
}
