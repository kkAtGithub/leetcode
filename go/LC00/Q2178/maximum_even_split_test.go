package Q2178

import (
	"testing"
)

func Test_maximumEvenSplit(t *testing.T) {
	type args struct {
		finalSum int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "1",
			args: args{
				finalSum: 28,
			},
			want: []int64{6, 4, 2, 16},
		},
		{
			name: "2",
			args: args{
				finalSum: 12,
			},
			want: []int64{6, 4, 2},
		},
		{
			name: "3",
			args: args{
				finalSum: 950,
			},
			want: []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 80},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumEvenSplit(tt.args.finalSum)

			if len(got) == len(tt.want) {
				var sum int64
				for _, i := range got {
					sum += i
				}
				if sum != tt.args.finalSum {
					t.Errorf("maximumEvenSplit() = %v, want %v", got, tt.want)
				}
				t.Logf("maximumEvenSplit() = %v, want %v", got, tt.want)
			} else {
				t.Errorf("maximumEvenSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
