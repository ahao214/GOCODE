package main

//541. 反转字符串 II
func reverseStr(s string, k int) string {
	sa := []byte(s)
	totalLength := len(sa)
	flag := true
	for i := 0; i < len(sa); i += k {
		if flag {
			left := i
			right := i + k - 1
			if right > totalLength-1 {
				right = totalLength - 1
			}
			reverseString(sa, left, right)
		}
		flag = !flag
	}
	return string(sa)
}

func reverseString(s []byte, left int, right int) {
	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp
		left++
		right--
	}
}
