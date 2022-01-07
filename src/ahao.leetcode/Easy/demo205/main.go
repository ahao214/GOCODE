package main

/*
输入：s = "egg", t = "add"
输出：true
*/

//205. 同构字符串
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return isIso(s, t) && isIso(t, s)
}

func isIso(s string, t string) bool {
	p := 0
	dic := make(map[byte]byte)
	for p < len(s)-1 {
		_, ok := dic[s[p]]
		if !ok {
			dic[s[p]] = t[p]
		} else {
			if dic[s[p]] != t[p] {
				return false
			}
		}
		p++
	}
	return true
}
