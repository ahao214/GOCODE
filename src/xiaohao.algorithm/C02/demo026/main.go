package main

//438. 找到字符串中所有字母异位词
//输入:s: "cbaebabacd" p: "abc"
//输出:[0, 6]
//解释: 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词
func findAnagrams(s string, p string) []int {
	return []int{}
}

func isSame(arr1 []int, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
