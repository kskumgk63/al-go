package implement_strStr

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "3",
			args: args{
				haystack: "aaaaa",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				haystack: "",
				needle:   "ss",
			},
			want: -1,
		},
		{
			name: "3",
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				haystack: "hello",
				needle:   "hello",
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				haystack: "hello",
				needle:   "e",
			},
			want: 1,
		},
		{
			name: "6",
			args: args{
				haystack: "hello",
				needle:   "ello",
			},
			want: 1,
		},
		{
			name: "7",
			args: args{
				haystack: "aaa",
				needle:   "aaaa",
			},
			want: -1,
		},
		{
			name: "8",
			args: args{
				haystack: "mississippi",
				needle:   "a",
			},
			want: -1,
		},
		{
			name: "9",
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
