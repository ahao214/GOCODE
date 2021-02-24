package main

//题目：无重复字符的最长子串
//说明：在⼀个字符串重寻找没有重复字⺟的最⻓⼦串
//输入：s = "abcabcbb"
//解释：因为无重复字符的最长子串是 "abc"，所以其长度为 3
//输出：3
//方法一：位图
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到
		//将X标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

//方法二：滑动窗口
func lengthOfLongestSubstringtwo(s string) int {
	if len(s) == 0 {
		return 0
	}
	var req [256]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && req[s[right+1]-'a'] == 0 {
			req[s[right+1]-'a']++
			right++
		} else {
			req[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
