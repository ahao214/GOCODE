package main

import "strings"

//290. 单词规律
func wordPattern(pattern string, str string) bool {
	strList := strings.Split(str, " ")
	patternByte := []byte(pattern)
	if pattern == "" || len(patternByte) != len(strList) {
		return false
	}

	pMap := map[byte]string{}
	sMap := map[string]byte{}
	for index, b := range patternByte {
		if _, ok := pMap[b]; !ok {
			if _, ok = sMap[strList[index]]; !ok {
				pMap[b] = strList[index]
				sMap[strList[index]] = b
			} else {
				if sMap[strList[index]] != b {
					return false
				}
			}
		} else {
			if pMap[b] != strList[index] {
				return false
			}
		}
	}
	return true
}

func wordPattern1(pattern string, s string) bool {
	patterns := strings.Split(pattern, "")
	strs := strings.Split(s, " ")
	if len(strs) != len(patterns) {
		return false
	}
	return checkPattern(strs, patterns) && checkPattern(patterns, strs)
}

func checkPattern(a []string, b []string) bool {
	dic := make(map[string]string)
	p := 0
	for p < len(a) {
		val, ok := dic[a[p]]
		if !ok {
			dic[a[p]] = b[p]
		} else {
			if val != b[p] {
				return false
			}
		}
		p++
	}
	return true
}
