# ECE654 Assignment 1: Static Analysis


## Usage
`go run cli.go -noIdentEqLen13 [filepath]`

`go run cli.go -maxControlNesting4 [filepath]`

##   `ast` library for Golang
The [`ast` library](https://pkg.go.dev/go/ast) belongs to the Standard library of the Go programming language. There are three important concepts in this library for static analysis: 
`func Walk(v Visitor, node Node)` for traversing the Abstract Syntax Tree, `type Visitor interface` for users to implement operations when visiting nodes , `type inspector func(Node) bool` expose the method for the user to implement those operations.



```go
// Walk traverses an AST in depth-first order: It starts by calling
// v.Visit(node); node must not be nil. If the visitor w returned by
// v.Visit(node) is not nil, Walk is invoked recursively with visitor
// w for each of the non-nil children of node, followed by a call of
// w.Visit(nil).
// Skeleton of 'Walk' function
func Walk(v Visitor, node Node) {
    if v = v.Visit(node); v == nil {
        return // base case
	}
	
	switch n := node.(type){
    // recursive visit every node 
		
                ......
		
		//example: 
		// For a declaration statement node,
		// Visit its only child "n.Decl"
		case *DeclStmt:
			Walk(v, n.Decl)
		......		
	}
	
	// this line of code could be superfluous
	v.Visit(nil)
	
}
```

```go
// A Visitor's Visit method is invoked for each node encountered by Walk.
// If the result visitor w is not nil, Walk visits each of the children
// of node with the visitor w, followed by a call of w.Visit(nil).
type Visitor interface {
	Visit(node Node) (w Visitor)
}
```


```go
type inspector func(Node) bool

func (f inspector) Visit(node Node) Visitor {
	if f(node) {
		return f
	}
	return nil
}

// Inspect traverses an AST in depth-first order: It starts by calling
// f(node); node must not be nil. If f returns true, Inspect invokes f
// recursively for each of the non-nil children of node, followed by a
// call of f(nil).
func Inspect(node Node, f func(Node) bool) {
	Walk(inspector(f), node)
}
```

The above implementation in the `ast` library provides an elegant way to check each node recursive. However, it seems over-complicated to implement an simple traversal feature with 2 interfaces. I have conjecture that it will still work after eliminating the `Visitor` class. The resulting API should be  more foolproof. 

While it is convenient to sort and optimize imports by using the API, The API exposed by this library is limited for performing more sophisticated static analysis than performing condition check on a single nodes. For example, counting nesting levels and type checking is impossible to implement with current API. The modern IDEs, such as GoVim and Goland, must have implemented their AST manipulation libraries. 

## My Implementation

### "There are no identifiers with length equal 13"
I implement the `Inspect` function exposed by `ast` library. I use a channel to collect the result. Once my code found a positive case, it will pass the result to the channel and stops the traversal. My test coverage for this analysis is thorough. I tested comments, variables, functions, type definitions, string literals, and even nested go files. 

### "Maximum control structure nesting of 4"
I refactored the `Inspector`, `Visitor` , `Walk` to pass a counter for backtracking the nesting level. My test coverage for this analysis is sketchy. I have tested for-loops, if-branching, and switch-branching. However, I skipped their combinations. 


## False Positives 
### "There are no identifiers with length equal 13"
It means that  while in reality a piece of code doesn't contain an identifier whose length equals 13, my system concludes that this piece of code does. I haven't found such cases.


### "Maximum control structure nesting of 4"
It means that  while in reality a piece of code doesn't contain a control structure that exceeded 4 levels of nesting, my system concludes that this piece of code does. I haven't found such cases.

## False Negative
### "There are no identifiers with length equal 13"
It means that  while in reality a piece of code does contain an identifier whose length equals 13, my system concludes that this piece of code doesn't. I haven't found such cases.


### "Maximum control structure nesting of 4"
It means that  while in reality a piece of code does contain a control structure that exceeded 4 levels of nesting, my system concludes that this piece of code doesn't. I haven't found such cases.
