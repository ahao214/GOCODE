package main

//784. 字母大小写全排列
func letterCasePermutation(s string) []string {
	res := []string{}
	if len(s) == 0 {
		return res
	}
	cur := []rune{}
	backtrack([]rune(s), 0, cur, &res)
	return res
}

func backtrack(str []rune, start int, cur []rune, res *[]string) {
	if len(str) == len(cur) {
		*res = append(*res, string(cur))
		return
	}
	for i := start; i < len(str); i++ {
		//add
		if !isLetter(str[i]) {
			cur = append(cur, str[i])
			backtrack(str, i+1, cur, res)
		} else {
			cur = append(cur, toUpperCase(str[i]))
			backtrack(str, i+1, cur, res)
			cur[len(cur)-1] = toLowerCase(str[i])
			backtrack(str, i+1, cur, res)
		}
		//remove
		cur = cur[:len(cur)-1]
	}
}

func isLetter(r rune) bool {
	return isLowerCase(r) || isUpperCase(r)
}

func isLowerCase(r rune) bool {
	return 97 <= r && r <= 122
}

func isUpperCase(r rune) bool {
	return 65 <= r && r <= 90
}

func toUpperCase(r rune) rune {
	if isUpperCase(r) {
		return r
	}
	return r - 32
}

func toLowerCase(r rune) rune {
	if isLowerCase(r) {
		return r
	}
	return r + 32
}
