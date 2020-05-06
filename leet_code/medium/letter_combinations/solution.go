package letter_combinations

//
// question: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
//

// cannot solve this question by myself
// refactored the best answer from discussion(https://leetcode.com/problems/letter-combinations-of-a-phone-number/discuss/195884/go-golang)
func letterCombinations(digits string) []string {
	if len(digits) < 1 || 4 < len(digits) {
		return []string{}
	}

	var letterMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var ret []string
	for _, r := range letterMap[string(digits[0])] {
		s := string(r)
		ret = append(ret, s)
	}
	for i := 1; i < len(digits); i++ {
		letter := letterMap[string(digits[i])]
		retCopy := ret   // var copy ret
		ret = []string{} // set as empty
		for j := range retCopy {
			for k := range letter {
				ret = append(ret, retCopy[j]+string(letter[k]))
			}
		}
	}

	return ret
}
