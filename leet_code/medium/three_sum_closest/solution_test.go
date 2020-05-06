package three_sum_closest

import (
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "(1)",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "(2)",
			args: args{
				nums:   []int{11, 2, 1, -4, 9},
				target: 9,
			},
			want: 9,
		},
		{
			name: "(3)",
			args: args{
				nums:   []int{},
				target: 1,
			},
			want: 0,
		},
		{
			name: "(4)",
			args: args{
				nums:   []int{0, 0, 0},
				target: 9,
			},
			want: 0,
		},
		{
			name: "(5)",
			args: args{
				nums:   []int{4, 5},
				target: 9,
			},
			want: 0,
		},
		{
			name: "(6)",
			args: args{
				nums:   []int{11, 11, 11, 11},
				target: 11,
			},
			want: 33,
		},
		{
			name: "(7)",
			args: args{
				nums:   []int{0, 2, 1, -3},
				target: 1,
			},
			want: 0,
		},
		{
			name: "(8)",
			args: args{
				nums:   []int{0, 2, 1, -3, 0, 0, 9, 2},
				target: -3,
			},
			want: -3,
		},
		{
			name: "(9)",
			args: args{
				nums:   []int{0, -2, -1, -3},
				target: -2,
			},
			want: -3,
		},
		{
			name: "(10)",
			args: args{
				nums:   []int{1, 2, 4, 8, 16, 32, 64, 128},
				target: 82,
			},
			want: 82,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
			if got := best_threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("best_threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
