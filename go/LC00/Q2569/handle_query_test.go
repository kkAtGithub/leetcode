package Q2569

import (
	"reflect"
	"testing"
)

func Test_handleQuery(t *testing.T) {
	type args struct {
		nums1   []int
		nums2   []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "1",
			args: args{
				nums1: []int{1, 0, 1},
				nums2: []int{0, 0, 0},
				queries: [][]int{
					{
						1, 1, 1,
					}, {
						2, 1, 0,
					}, {
						3, 0, 0,
					},
				},
			},
			want: []int64{
				3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleQuery(tt.args.nums1, tt.args.nums2, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
