package main

/*
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

输入：strs = ["flower","flow","flight"]
输出："fl"
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		resLength := len(res)
		tmpLength := len(strs[i])
		length := min(resLength, tmpLength)
		j := 0
		for j < length {
			if res[j] != strs[i][j] {
				break
			}
			j++
		}
		res = res[:j]
		if len(res) == 0 {
			return ""
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
