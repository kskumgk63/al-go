package implement_strStr

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	searchLimit := len(needle)
	if searchLimit > haystackLen {
		return -1
	}
	for i := 0; i <= haystackLen; i++ {
		if i+searchLimit > haystackLen {
			break
		}
		if haystack[i:i+searchLimit] == needle {
			return i
		}
	}

	return -1
	/*
		there is uncompleted method blow.
		failed to run on test case "6".

			if needle == "" || needle == haystack {
				return 0
			}
			if haystack == "" || len(haystack) < len(needle) {
				return -1
			}

			var (
				count     = 0
				idxI      = 0
				idxII     = 0
				resultIdx = -1
			)
			for i := idxI; i < len(needle); i++ {
				n := string(needle[i])
				for j := idxII; j < len(haystack); j++ {
					h := string(haystack[j])
					if n == h {
						if count == 0 {
							if i == len(needle)-1 && len(needle) != 1 {
								break
							}
							resultIdx = j
						}
						count++
						idxII = j + 1
						break
					}
				}
				if 0 < count && count < len(needle) {
					count = 0
					idxI = 0
				}
			}
			return resultIdx
	*/
}
