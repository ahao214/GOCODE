package main

/*
467. 环绕字符串中唯一的子字符串
把字符串 s 看作 "abcdefghijklmnopqrstuvwxyz" 的无限环绕字符串，所以 s 看起来是这样的：

"...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd...." 。
现在给定另一个字符串 p 。返回 s 中 不同 的 p 的 非空子串 的数量 。

示例 1：

输入：p = "a"
输出：1
解释：字符串 s 中只有 p 的一个 "a" 子字符。
示例 2：

输入：p = "cac"
输出：2
解释：字符串 s 中只有 p 的两个子串 ("a", "c") 。
示例 3：

输入：p = "zab"
输出：6
解释：在字符串 s 中有 p 的六个子串 ("z", "a", "b", "za", "ab", "zab") 。


提示：

1 <= p.length <= 105
p 由小写英文字母组成

*/
func findSubstringInWraproundString(p string) (ans int) {
	dp := [26]int{}
	k := 0
	for i, ch := range p {
		if i > 0 && (byte(ch)-p[i-1]+26)%26 == 1 { // 字符之差为 1 或 -25
			k++
		} else {
			k = 1
		}
		dp[ch-'a'] = max(dp[ch-'a'], k)
	}
	for _, v := range dp {
		ans += v
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
