package test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Test() string {
	// Define the BASIC program to be translated
	basicCode := `
	package main

	func main(){
		fmt.Println("Hello World")
	}
`

	// Create a new file set
	fset := token.NewFileSet()

	// Parse the BASIC code and generate the AST
	file, err := parser.ParseFile(fset, "", basicCode, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing BASIC code:", err)
		return ""
	}

	// Traverse the AST and generate C code
	cCode := ""
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.CallExpr:
			if ident, ok := node.Fun.(*ast.Ident); ok {
				switch ident.Name {
				case "PRINT":
					if len(node.Args) > 0 {
						arg := node.Args[0]
						if basicStr, ok := arg.(*ast.BasicLit); ok && basicStr.Kind == token.STRING {
							cCode += fmt.Sprintf("printf(%s);\n", basicStr.Value)
						}
					}
				case "END":
					cCode += "return 0;\n"
				}
			}
		}
		return true
	})

	// Print the generated C code
	fmt.Println(cCode)
	return cCode
}
