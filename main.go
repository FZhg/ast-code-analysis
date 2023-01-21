package main

import (
	"flag"
	"fmt"
)

var noIdentEqLen13 = flag.String("noIdentEqLen13", "", "file path for checking if any identifier's length in the file is equal to 13")
var maxControlNesting4 = flag.String("maxControlNesting4", "", "file path for checking if any control structure in the file exceeds 4 nesting level")

func main() {
	flag.Parse()
	if *noIdentEqLen13 != "" {
		node := getRootNode(*noIdentEqLen13)
		fmt.Printf("%t\n", NoIdentifierLenEqual13(node))
	}

	if *maxControlNesting4 != "" {
		node := getRootNode(*maxControlNesting4)
		fmt.Printf("%t\n", ExceedMaxNestingMax4(node))
	}
}
