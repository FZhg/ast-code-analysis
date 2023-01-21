package ast_code_analysis

import (
	"go/ast"
	"testing"
)

func Test_isMaxNestingMax4(t *testing.T) {
	type args struct {
		node ast.Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "for1",
			args: args{
				node: getRootNode("max_control_nesting_4_testcases/for1.txt"),
			},
			want: false,
		},
		{
			name: "for4",
			args: args{
				node: getRootNode("max_control_nesting_4_testcases/for4.txt"),
			},
			want: false,
		},
		{
			name: "for5",
			args: args{
				node: getRootNode("max_control_nesting_4_testcases/for5.txt"),
			},
			want: true,
		},
		{
			name: "if1",
			args: args{
				node: getRootNode("max_control_nesting_4_testcases/if1.txt"),
			},
			want: false,
		},
		{
			name: "if-(else-if)-if-else-if3",
			args: args{
				node: getRootNode("max_control_nesting_4_testcases/if-(else-if)-if-else-if3.txt"),
			},
			want: false,
		},
		{
			name: "if-(else-if)-if-else-if-else-if-if-if5",
			args: args{
				node: getRootNode("max_control_nesting_4_testcases/if-(else-if)-if-else-if-else-if-if-if5.txt"),
			},
			want: true,
		},
		{
			name: "switch-if-switch-if-for5.txt",
			args: args{
				node: getRootNode("max_control_nesting_4_testcases/switch-if-switch-if-for5.txt"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceedMaxNestingMax4(tt.args.node); got != tt.want {
				t.Errorf("max_control_nesting_4_testcases() = %v, want %v", got, tt.want)
			}
		})
	}
}
