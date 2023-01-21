package ast_static_analysis

import (
	"go/ast"
)

func ExceedMaxNestingMax4(node ast.Node) bool {
	nestingCounter := make(chan int, 1)
	// Inspect will update the counter and traverse the tree if func argument returns true.
	Inspect(node, 0, func(node2 ast.Node, counter int) (bool, int) {

		if len(nestingCounter) == 1 {
			return false, counter // found the case; stop the traversal of siblings
		}
		if counter > 4 {
			nestingCounter <- counter
			return false, counter // found the case; stop the traversal of children
		}

		switch node2.(type) {
		// There are 5 control structure in Golang: For, if, select, switch, type-switch
		// else-if, range-for, select-case, switch-case, type-switch-case are on the same nested level.
		case *ast.ForStmt, *ast.IfStmt, *ast.SelectStmt, *ast.SwitchStmt, *ast.TypeSwitchStmt:
			counter++
		}
		return true, counter
	})

	if len(nestingCounter) == 0 {
		return false
	} else {
		return true
	}
}
