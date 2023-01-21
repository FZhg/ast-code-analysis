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
				node: getRootNode("no_ident_equal_len_13_testcases/singleVariableLen13.txt"),
			},
			want: false,
		},
		{
			name: "noVariableLen13",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/noVariableLen13.txt"),
			},
			want: true,
		},
		{
			name: "threeVariablesLen13",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/threeVariablesLen13.txt"),
			},
			want: false,
		},
		{
			name: "singleFuncLen13",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/singleFuncLen13.txt"),
			},
			want: false,
		},
		{
			name: "twoMethodsLen13",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/twoFuncLen13.txt"),
			},
			want: false,
		},
		{
			name: "noFuncLen13",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/noFuncLen13.txt"),
			},
			want: true,
		},
		{
			name: "singleStrLen13",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/singleStrLen13.txt"),
			},
			want: true,
		},
		{
			name: "noStrLen13",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/noStrLen13.txt"),
			},
			want: true,
		},
		{
			name: "noCommentsLen13",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/noFuncLen13.txt"),
			},
			want: true,
		},
		{
			name: "singleCommentLen13",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/singleCommentLen13.txt"),
			},
			want: true,
		},
		{
			name: "typeDefLen13",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/typeDefLen13.txt"),
			},
			want: false,
		},
		{
			name: "nestedGoFilesNegative",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/nestedGoFilesNegative.txt"),
			},
			want: true,
		},
		{
			name: "nestedGoFilesPositive",
			args: args{
				node: getRootNode("no_ident_equal_len_13_testcases/nestedGoFilesPositive.txt"),
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
