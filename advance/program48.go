package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
)

func main() {
	src := `package main
func main() {
	oldVar := 42
}`

	fset := token.NewFileSet()
	node, _ := parser.ParseFile(fset, "", src, 0)

	// Walk the AST and modify the source code
	ast.Inspect(node, func(n ast.Node) bool {
		ident, ok := n.(*ast.Ident)
		// Find any variable named 'oldVar' and rename it to 'newSuperVar'
		if ok && ident.Name == "oldVar" {
			ident.Name = "newSuperVar"
		}
		return true
	})

	// Print the modified AST back into valid Go source code
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, node)

	fmt.Println("--- Rewritten Code ---")
	fmt.Println(buf.String())
}