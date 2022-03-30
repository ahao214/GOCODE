package main

/*
389. 找不同
给定两个字符串 s 和 t，它们只包含小写字母。
字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
请找出在 t 中被添加的字母。
*/
func findTheDifference(s string, t string) byte {
	var res byte
	if len(s) > len(t) {
		return res
	}
	T := []byte(t)
	S := []byte(s)
	dic := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dic[S[i]]++
	}
	for i := 0; i < len(T); i++ {
		val, ok := dic[T[i]]
		if !ok {
			res = T[i]
			break
		} else {
			if val == 0 {
				res = T[i]
				break
			}
			dic[T[i]]--
		}
	}
	return res
}
