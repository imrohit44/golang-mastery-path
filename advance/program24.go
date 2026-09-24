package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `
package math
func Add(a, b int) int { return a + b }
func Subtract(a, b int) int { return a - b }
`
	fset := token.NewFileSet()
	// Parse the string into an AST
	node, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("Functions found in source code:")
	
	// Walk the AST looking for function declarations
	ast.Inspect(node, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			fmt.Println("-", fn.Name.Name)
		}
		return true
	})
}