package main

//1374. 生成每种字符都是奇数个的字符串
func generateTheString(n int) string {
	if n == 0 {
		return ""
	}
	if n%2 == 0 {
		return buildstring(n-1) + "b"
	} else {
		return buildstring(n)
	}
}

func buildstring(n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += "a"
	}
	return res
}
