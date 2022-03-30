package main

//344. 反转字符串
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseStr(s []byte) {
	if s == nil || len(s) == 0 {
		return
	}
	left, right := 0, len(s)-1
	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp
		left++
		right--
	}
}
