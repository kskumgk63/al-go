package dynamic_programming

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n    int
		memo map[int]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 0",
			args: args{
				n:    0,
				memo: make(map[int]int),
			},
			want: 0,
		},
		{
			name: "n = 1",
			args: args{
				n:    1,
				memo: make(map[int]int),
			},
			want: 1,
		},
		{
			name: "n = 2",
			args: args{
				n:    2,
				memo: make(map[int]int),
			},
			want: 1,
		},
		{
			name: "n = 3",
			args: args{
				n:    3,
				memo: make(map[int]int),
			},
			want: 2,
		},
		{
			name: "n = 4",
			args: args{
				n:    4,
				memo: make(map[int]int),
			},
			want: 3,
		},
		{
			name: "n = 5",
			args: args{
				n:    5,
				memo: make(map[int]int),
			},
			want: 5,
		},
		{
			name: "n = 6",
			args: args{
				n:    6,
				memo: make(map[int]int),
			},
			want: 8,
		},
		{
			name: "n = 10",
			args: args{
				n:    10,
				memo: make(map[int]int),
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n, tt.args.memo); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
