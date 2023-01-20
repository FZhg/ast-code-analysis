package ast_code_analysis

import (
	"go/ast"
	"testing"
)

func TestNoIdentifierLenEqual13(t *testing.T) {
	type args struct {
		node ast.Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "singleVariableLen13",
			args: args{
				node: getRootNode("testcases/singleVariableLen13.txt"),
			},
			want: false,
		},
		{
			name: "noVariableLen13",
			args: args{
				node: getRootNode("testcases/noVariableLen13.txt"),
			},
			want: true,
		},
		{
			name: "threeVariablesLen13",
			args: args{
				node: getRootNode("testcases/threeVariablesLen13.txt"),
			},
			want: false,
		},
		{
			name: "nestedGoFilesNegative",
			args: args{
				node: getRootNode("testcases/nestedGoFilesNegative.txt"),
			},
			want: true,
		},
		{
			name: "nestedGoFilesPositive",
			args: args{
				node: getRootNode("testcases/nestedGoFilesPositive.txt"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoIdentifierLenEqual13(tt.args.node); got != tt.want {
				t.Errorf("NoIdentifierLenEqual13() = %v, want %v", got, tt.want)
			}
		})
	}
}