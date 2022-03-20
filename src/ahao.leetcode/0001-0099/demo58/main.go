package main

/*
58. 最后一个单词的长度
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。
单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
*/
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	first := -1
	chars := []byte(s)
	for i := len(chars) - 1; i >= 0; i-- {
		if first == -1 {
			if chars[i] == ' ' {
				continue
			} else {
				first = i
			}
		} else {
			if chars[i] == ' ' {
				return first - i
			}
		}
	}
	return first + 1
}
