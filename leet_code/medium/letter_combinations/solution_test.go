package letter_combinations

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "(8)",
			args: args{
				digits: "999",
			},
			want: []string{"www", "wwx", "wwy", "wwz", "wxw", "wxx", "wxy", "wxz", "wyw", "wyx", "wyy", "wyz", "wzw", "wzx", "wzy", "wzz", "xww", "xwx", "xwy", "xwz", "xxw", "xxx", "xxy", "xxz", "xyw", "xyx", "xyy", "xyz", "xzw", "xzx", "xzy", "xzz", "yww", "ywx", "ywy", "ywz", "yxw", "yxx", "yxy", "yxz", "yyw", "yyx", "yyy", "yyz", "yzw", "yzx", "yzy", "yzz", "zww", "zwx", "zwy", "zwz", "zxw", "zxx", "zxy", "zxz", "zyw", "zyx", "zyy", "zyz", "zzw", "zzx", "zzy", "zzz"},
		},
		{
			name: "(7)",
			args: args{
				digits: "8632",
			},
			want: []string{"tmda", "tmdb", "tmdc", "tmea", "tmeb", "tmec", "tmfa", "tmfb", "tmfc", "tnda", "tndb", "tndc", "tnea", "tneb", "tnec", "tnfa", "tnfb", "tnfc", "toda", "todb", "todc", "toea", "toeb", "toec", "tofa", "tofb", "tofc", "umda", "umdb", "umdc", "umea", "umeb", "umec", "umfa", "umfb", "umfc", "unda", "undb", "undc", "unea", "uneb", "unec", "unfa", "unfb", "unfc", "uoda", "uodb", "uodc", "uoea", "uoeb", "uoec", "uofa", "uofb", "uofc", "vmda", "vmdb", "vmdc", "vmea", "vmeb", "vmec", "vmfa", "vmfb", "vmfc", "vnda", "vndb", "vndc", "vnea", "vneb", "vnec", "vnfa", "vnfb", "vnfc", "voda", "vodb", "vodc", "voea", "voeb", "voec", "vofa", "vofb", "vofc"},
		},
		{
			name: "(6)",
			args: args{
				digits: "86321",
			},
			want: []string{},
		},
		{
			name: "(5)",
			args: args{
				digits: "444",
			},
			want: []string{"ggg", "ggh", "ggi", "ghg", "ghh", "ghi", "gig", "gih", "gii", "hgg", "hgh", "hgi", "hhg", "hhh", "hhi", "hig", "hih", "hii", "igg", "igh", "igi", "ihg", "ihh", "ihi", "iig", "iih", "iii"},
		},
		{
			name: "(4)",
			args: args{
				digits: "98",
			},
			want: []string{"wt", "wu", "wv", "xt", "xu", "xv", "yt", "yu", "yv", "zt", "zu", "zv"},
		},
		{
			name: "(3)",
			args: args{
				digits: "0721",
			},
			want: []string{},
		},
		{
			name: "(2)",
			args: args{
				digits: "",
			},
			want: []string{},
		},
		{
			name: "(1)",
			args: args{
				digits: "23",
			},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.args.digits)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Run() differs: (-got +want)\n%s", diff)
			}
		})
	}
}
