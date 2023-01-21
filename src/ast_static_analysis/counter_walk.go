// I refactored (Visitor, Walker, Inspector API of the ast packages)[https://cs.opensource.google/go/go/+/refs/tags/go1.19.5:src/go/ast/walk.go;l=50].
// The official package is limited  because user can not keep a counter for the appearance
// of a certain conditions but can check if some condition occurs

package ast_static_analysis

import (
	"fmt"
	"go/ast"
)

// Visitor's Visit method is invoked for each node encountered by Walk.
// If the result visitor w is not nil, Walk visits each of the children
// of node with the visitor w, followed by a call of w.Visit(nil).
type Visitor interface {
	Visit(node ast.Node, parentCounter int) (w Visitor, childCounter int)
}

// Helper functions for common node lists. They may be empty.

func walkIdentList(v Visitor, list []*ast.Ident, counter int) {
	for _, x := range list {
		Walk(v, x, counter)
	}
}

func walkExprList(v Visitor, list []ast.Expr, counter int) {
	for _, x := range list {
		Walk(v, x, counter)
	}
}

func walkStmtList(v Visitor, list []ast.Stmt, counter int) {
	for _, x := range list {
		Walk(v, x, counter)
	}
}

func walkDeclList(v Visitor, list []ast.Decl, counter int) {
	for _, x := range list {
		Walk(v, x, counter)
	}
}

// TODO(gri): Investigate if providing a closure to Walk leads to
// simpler use (and may help eliminate Inspect in turn).

// Walk traverses an AST in depth-first order: It starts by calling
// v.Visit(node); node must not be nil. If the visitor w returned by
// v.Visit(node) is not nil, Walk is invoked recursively with visitor
// w for each of the non-nil children of node, followed by a call of
// w.Visit(nil).
func Walk(v Visitor, node ast.Node, counter int) {
	v, counter = v.Visit(node, counter)
	if v == nil {
		return
	}

	// walk children
	// (the order of the cases matches the order
	// of the corresponding node types in ast.go)
	switch n := node.(type) {
	// Comments and fields
	case *ast.Comment:
		// nothing to do

	case *ast.CommentGroup:
		for _, c := range n.List {
			Walk(v, c, counter)
		}

	case *ast.Field:
		if n.Doc != nil {
			Walk(v, n.Doc, counter)
		}
		walkIdentList(v, n.Names, counter)
		if n.Type != nil {
			Walk(v, n.Type, counter)
		}
		if n.Tag != nil {
			Walk(v, n.Tag, counter)
		}
		if n.Comment != nil {
			Walk(v, n.Comment, counter)
		}

	case *ast.FieldList:
		for _, f := range n.List {
			Walk(v, f, counter)
		}

	// Expressions
	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
		// nothing to do

	case *ast.Ellipsis:
		if n.Elt != nil {
			Walk(v, n.Elt, counter)
		}

	case *ast.FuncLit:
		Walk(v, n.Type, counter)
		Walk(v, n.Body, counter)

	case *ast.CompositeLit:
		if n.Type != nil {
			Walk(v, n.Type, counter)
		}
		walkExprList(v, n.Elts, counter)

	case *ast.ParenExpr:
		Walk(v, n.X, counter)

	case *ast.SelectorExpr:
		Walk(v, n.X, counter)
		Walk(v, n.Sel, counter)

	case *ast.IndexExpr:
		Walk(v, n.X, counter)
		Walk(v, n.Index, counter)

	case *ast.IndexListExpr:
		Walk(v, n.X, counter)
		for _, index := range n.Indices {
			Walk(v, index, counter)
		}

	case *ast.SliceExpr:
		Walk(v, n.X, counter)
		if n.Low != nil {
			Walk(v, n.Low, counter)
		}
		if n.High != nil {
			Walk(v, n.High, counter)
		}
		if n.Max != nil {
			Walk(v, n.Max, counter)
		}

	case *ast.TypeAssertExpr:
		Walk(v, n.X, counter)
		if n.Type != nil {
			Walk(v, n.Type, counter)
		}

	case *ast.CallExpr:
		Walk(v, n.Fun, counter)
		walkExprList(v, n.Args, counter)

	case *ast.StarExpr:
		Walk(v, n.X, counter)

	case *ast.UnaryExpr:
		Walk(v, n.X, counter)

	case *ast.BinaryExpr:
		Walk(v, n.X, counter)
		Walk(v, n.Y, counter)

	case *ast.KeyValueExpr:
		Walk(v, n.Key, counter)
		Walk(v, n.Value, counter)

	// Types
	case *ast.ArrayType:
		if n.Len != nil {
			Walk(v, n.Len, counter)
		}
		Walk(v, n.Elt, counter)

	case *ast.StructType:
		Walk(v, n.Fields, counter)

	case *ast.FuncType:
		if n.TypeParams != nil {
			Walk(v, n.TypeParams, counter)
		}
		if n.Params != nil {
			Walk(v, n.Params, counter)
		}
		if n.Results != nil {
			Walk(v, n.Results, counter)
		}

	case *ast.InterfaceType:
		Walk(v, n.Methods, counter)

	case *ast.MapType:
		Walk(v, n.Key, counter)
		Walk(v, n.Value, counter)

	case *ast.ChanType:
		Walk(v, n.Value, counter)

	// Statements
	case *ast.BadStmt:
		// nothing to do

	case *ast.DeclStmt:
		Walk(v, n.Decl, counter)

	case *ast.EmptyStmt:
		// nothing to do

	case *ast.LabeledStmt:
		Walk(v, n.Label, counter)
		Walk(v, n.Stmt, counter)

	case *ast.ExprStmt:
		Walk(v, n.X, counter)

	case *ast.SendStmt:
		Walk(v, n.Chan, counter)
		Walk(v, n.Value, counter)

	case *ast.IncDecStmt:
		Walk(v, n.X, counter)

	case *ast.AssignStmt:
		walkExprList(v, n.Lhs, counter)
		walkExprList(v, n.Rhs, counter)

	case *ast.GoStmt:
		Walk(v, n.Call, counter)

	case *ast.DeferStmt:
		Walk(v, n.Call, counter)

	case *ast.ReturnStmt:
		walkExprList(v, n.Results, counter)

	case *ast.BranchStmt:
		if n.Label != nil {
			Walk(v, n.Label, counter)
		}

	case *ast.BlockStmt:
		walkStmtList(v, n.List, counter)

	// control structure 1
	case *ast.IfStmt:
		if n.Init != nil {
			Walk(v, n.Init, counter)
		}
		Walk(v, n.Cond, counter)
		Walk(v, n.Body, counter)
		if n.Else != nil {
			Walk(v, n.Else, counter)
		}

	case *ast.CaseClause:
		walkExprList(v, n.List, counter)
		walkStmtList(v, n.Body, counter)

	// control structure 2
	case *ast.SwitchStmt:
		if n.Init != nil {
			Walk(v, n.Init, counter)
		}
		if n.Tag != nil {
			Walk(v, n.Tag, counter)
		}
		Walk(v, n.Body, counter)

	// control structure 3
	case *ast.TypeSwitchStmt:
		if n.Init != nil {
			Walk(v, n.Init, counter)
		}
		Walk(v, n.Assign, counter)
		Walk(v, n.Body, counter)

	case *ast.CommClause:
		if n.Comm != nil {
			Walk(v, n.Comm, counter)
		}
		walkStmtList(v, n.Body, counter)

	// control structure 4
	case *ast.SelectStmt:
		Walk(v, n.Body, counter)

	// control structure 5
	case *ast.ForStmt:
		if n.Init != nil {
			Walk(v, n.Init, counter)
		}
		if n.Cond != nil {
			Walk(v, n.Cond, counter)
		}
		if n.Post != nil {
			Walk(v, n.Post, counter)
		}
		Walk(v, n.Body, counter)

	case *ast.RangeStmt:
		if n.Key != nil {
			Walk(v, n.Key, counter)
		}
		if n.Value != nil {
			Walk(v, n.Value, counter)
		}
		Walk(v, n.X, counter)
		Walk(v, n.Body, counter)

	// Declarations
	case *ast.ImportSpec:
		if n.Doc != nil {
			Walk(v, n.Doc, counter)
		}
		if n.Name != nil {
			Walk(v, n.Name, counter)
		}
		Walk(v, n.Path, counter)
		if n.Comment != nil {
			Walk(v, n.Comment, counter)
		}

	case *ast.ValueSpec:
		if n.Doc != nil {
			Walk(v, n.Doc, counter)
		}
		walkIdentList(v, n.Names, counter)
		if n.Type != nil {
			Walk(v, n.Type, counter)
		}
		walkExprList(v, n.Values, counter)
		if n.Comment != nil {
			Walk(v, n.Comment, counter)
		}

	case *ast.TypeSpec:
		if n.Doc != nil {
			Walk(v, n.Doc, counter)
		}
		Walk(v, n.Name, counter)
		if n.TypeParams != nil {
			Walk(v, n.TypeParams, counter)
		}
		Walk(v, n.Type, counter)
		if n.Comment != nil {
			Walk(v, n.Comment, counter)
		}

	case *ast.BadDecl:
		// nothing to do

	case *ast.GenDecl:
		if n.Doc != nil {
			Walk(v, n.Doc, counter)
		}
		for _, s := range n.Specs {
			Walk(v, s, counter)
		}

	case *ast.FuncDecl:
		if n.Doc != nil {
			Walk(v, n.Doc, counter)
		}
		if n.Recv != nil {
			Walk(v, n.Recv, counter)
		}
		Walk(v, n.Name, counter)
		Walk(v, n.Type, counter)
		if n.Body != nil {
			Walk(v, n.Body, counter)
		}

	// Files and packages
	case *ast.File:
		if n.Doc != nil {
			Walk(v, n.Doc, counter)
		}
		Walk(v, n.Name, counter)
		walkDeclList(v, n.Decls, counter)
		// don't walk n.Comments - they have been
		// visited already through the individual
		// nodes

	case *ast.Package:
		for _, f := range n.Files {
			Walk(v, f, counter)
		}

	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}

	v.Visit(nil, counter) // this line should not be included
}

type inspector func(ast.Node, int) (bool, int)

func (f inspector) Visit(node ast.Node, counter int) (Visitor, int) {
	shouldWalk, counter := f(node, counter)
	if shouldWalk {
		return f, counter
	}
	return nil, counter
}

// Inspect traverses an AST in depth-first order: It starts by calling
// f(node); node must not be nil. If f returns true, Inspect invokes f
// recursively for each of the non-nil children of node, followed by a
// call of f(nil).
func Inspect(node ast.Node, counter int, f func(node ast.Node, counter int) (bool, int)) {
	Walk(inspector(f), node, counter)
}
