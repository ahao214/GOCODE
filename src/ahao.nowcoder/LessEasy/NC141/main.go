package main

//判断回文

/**
 * @param str string字符串 待判断的字符串
 * @return bool布尔型
 */
func judge(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
