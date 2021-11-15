package main

//NC103 反转字符串
func solve(str string) string {
	if len(str) == 0 {
		return ""
	}
	sByte := []byte(str)
	left := 0
	right := len(sByte) - 1
	for left < right {
		tmp := sByte[left]
		sByte[left] = sByte[right]
		sByte[right] = tmp
		left++
		right--
	}
	return string(sByte)
}
