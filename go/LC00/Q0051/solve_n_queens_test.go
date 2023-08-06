package Q0051

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1",
			args: args{
				n: 4,
			},
			want: [][]string{
				{
					".Q..",
					"...Q",
					"Q...",
					"..Q.",
				},
				{
					"..Q.",
					"Q...",
					"...Q",
					".Q..",
				},
			},
		},
		{
			name: "2",
			args: args{
				n: 8,
			},
			want: [][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Log(len(got))
				for _, board := range got {
					for _, row := range board {
						t.Log(row)
					}
					t.Log("")
				}
			}
		})
	}
}
