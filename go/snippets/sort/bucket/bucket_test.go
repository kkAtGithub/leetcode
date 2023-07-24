package bucket

import (
	"sort"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		nums sort.IntSlice
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums: []int{0, -1, 3, -4},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{0, -1, 3, -4, 7, 6, 5},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{0, -1, 3, -4, 7, 11, 9, -2, 39, -23, -12, -43, 398},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortSlice(tt.args.nums)
			t.Logf("nums after sort, %v", tt.args.nums)
			if !sort.IsSorted(tt.args.nums) {
				t.Errorf("nums is not sorted, %v", tt.args.nums)
			}
		})
	}
}
