package main

import "strings"

//709. 转换成小写字母
func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func toLowerCase2(s string) string {
	lower := &strings.Builder{}
	lower.Grow(len(s))
	for _, ch := range s {
		if 65 <= ch && ch <= 90 {
			ch |= 32
		}
		lower.WriteRune(ch)
	}
	return lower.String()
}
