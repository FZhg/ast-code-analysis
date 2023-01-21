package main

import (
	"assignment1ECE654/src/ast_static_analysis"
	"flag"
	"fmt"
)

var noIdentEqLen13 = flag.String("noIdentEqLen13", "", "file path for checking if any identifier's length in the file is equal to 13")
var maxControlNesting4 = flag.String("maxControlNesting4", "", "file path for checking if any control structure in the file exceeds 4 nesting level")

func main() {
	flag.Parse()
	if *noIdentEqLen13 != "" {
		node := ast_static_analysis.GetRootNode(*noIdentEqLen13)
		fmt.Printf("%t\n", ast_static_analysis.NoIdentifierLenEqual13(node))
	}

	if *maxControlNesting4 != "" {
		node := ast_static_analysis.GetRootNode(*maxControlNesting4)
		fmt.Printf("%t\n", ast_static_analysis.ExceedMaxNestingMax4(node))
	}
}
