package main

import "strings"

//500. 键盘行
func findWords(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	res := make([]string, 0)
	for _, s := range words {
		if len(s) == 0 {
			continue
		}
		lowerS := strings.ToLower(s)
		sameRow := false
		for _, r := range rows {
			//判断r中是否包含lowerS中任何一个字符
			if strings.ContainsAny(lowerS, r) {
				sameRow = !sameRow
				if !sameRow {
					break
				}
			}
		}
		if sameRow {
			res = append(res, s)
		}
	}
	return res
}
