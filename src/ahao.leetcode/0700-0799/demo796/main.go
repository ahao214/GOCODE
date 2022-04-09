package main

/*
796. 旋转字符串
给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。

s 的 旋转操作 就是将 s 最左边的字符移动到最右边。

例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。
*/
func rotateString(s, goal string) bool {
	n := len(s)
	if n != len(goal) {
		return false
	}
next:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[(i+j)%n] != goal[j] {
				continue next
			}
		}
		return true
	}
	return false
}
