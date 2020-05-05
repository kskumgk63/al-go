package longest_common_prefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "(1)",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "(2)",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "(3)",
			args: args{
				strs: []string{""},
			},
			want: "",
		},
		{
			name: "(4)",
			args: args{
				strs: []string{"hello", "hello", "hello"},
			},
			want: "hello",
		},
		{
			name: "(5)",
			args: args{
				strs: []string{"mississippi", "missmiss", "misii"},
			},
			want: "mis",
		},
		{
			name: "(6)",
			args: args{
				strs: []string{"mississippi", "", "misii"},
			},
			want: "",
		},
		{
			name: "(7)",
			args: args{
				strs: []string{"aprefix", "bprefix", "cprefix"},
			},
			want: "",
		},
		{
			name: "(8)",
			args: args{
				strs: []string{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
