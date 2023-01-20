package ast_code_analysis

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// NoIdentifierLenEqual13 returns true if no identifier's
// length equals 13. returns false otherwise.
// ast.Node is an interface. ast.File are legitimate input
// since it implements the interface.
func NoIdentifierLenEqual13(node ast.Node) bool {
	// the ast.inspect function will keep invoke function argument recursive on the children
	// nodes if function argument returns true
	result := make(chan bool, 1)
	ast.Inspect(node, func(node ast.Node) bool {
		if len(result) == 1 {
			return false // stop traversal because
			// this function already found an positive case
		}

		iden, ok := node.(*ast.Ident) // use reflection to cast node to specific identifier node
		// what is an identifier? a name for method, struct or any other user defined stuff
		if ok && (len(iden.Name) == 13) {
			result <- false
			return false
		}
		return true
	})

	fmt.Printf("Result Length: %d\n", len(result))
	if len(result) == 0 {
		return true // this method have looped over entire tree
		// and found no such cases
	} else {
		return false
	}
}

// Create root node of ast from the filename
func getRootNode(filename string) ast.Node {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		panic(err)
	}
	return node
}
