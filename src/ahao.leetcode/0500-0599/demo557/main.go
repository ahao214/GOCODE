package main

import "strings"

//557. 反转字符串中的单词 III
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, w := range words {
		words[i] = reverseString(w)
	}
	return strings.Join(words, " ")
}

//反转每一个单词
func reverseString(word string) string {
	s := []byte(word)
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return string(s)
}
