package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	base := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(base) {
			base = strs[i]
		}
	}
	ret := ""
	for offset := 1; offset <= len(base); offset++ {
		isMatch := true
		for _, str := range strs {
			if base[0:offset] != str[0:offset] {
				isMatch = false
			}
		}
		if isMatch && len(ret) < offset {
			ret = base[0:offset]
		}
		if 0 == len(ret) {
			break
		}
	}
	return ret
}
