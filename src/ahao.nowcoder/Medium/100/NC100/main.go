package main

//NC100 把字符串转换成整数(atoi)
func StrToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	var int32_max = 2147483647
	var int32_min = -2147483648
	var sign = 1
	var i int
	// 跳过字符头部数值
	for i < len(s) && s[i] == ' ' {
		i++
	}
	// 如果都是空格，则返回0
	if i == len(s) {
		return 0
	}
	// 如果是符号位则i++
	if s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	var res int
	// 开始处理数值
	for i < len(s) {
		// 不符合规则则跳过
		if s[i] < '0' || s[i] > '9' {
			break
		}
		// 处理越界的情况
		res = res*10 + int(s[i]-'0')
		if res > int32_max {
			if sign == 1 {
				return int32_max
			} else {
				return int32_min
			}
		}
		i++
	}
	return sign * res
}
