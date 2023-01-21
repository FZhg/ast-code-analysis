package ast_static_analysis

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// NoIdentifierLenEqual13 returns true if no identifier's
// length equals 13. returns false otherwise.
// ast.Node is an interface. ast.File are legitimate input
// since it implements the interface.
// An identifier is anything other than a keyword or a literal (strings, int ....).
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
		if ok && (len(iden.Name) == 13) {
			result <- false
			return false
		}
		return true
	})
	if len(result) == 0 {
		return true // this method have looped over entire tree
		// and found no such cases
	} else {
		return false
	}
}

// GetRootNode Create root node of ast from the filename
func GetRootNode(filename string) ast.Node {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		panic(err)
	}
	return node
}
