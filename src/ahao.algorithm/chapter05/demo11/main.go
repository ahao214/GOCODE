package main

import (
	"fmt"
)

//如何按照给定的字母序列对字符数组排序
func compare(str1, str2 string, charToInt map[byte]int) int {
	len1, len2 := len(str1), len(str2)
	i, j := 0, 0
	for i < len1 && j < len2 {
		//如果字符不在给定的序列中，把值赋为-1
		if _, ok := charToInt[str1[i]]; !ok {
			charToInt[str1[i]] = -1
		}
		if _, ok := charToInt[str2[j]]; !ok {
			charToInt[str2[j]] = -1
		}
		//比较各个字符的大小
		if charToInt[str1[i]] < charToInt[str2[j]] {
			return -1
		} else if charToInt[str1[i]] > charToInt[str2[j]] {
			return 1
		} else {
			i++
			j++
		}
	}
	if i == len1 && j == len2 {
		return 0
	} else if i == len1 {
		return -1
	} else {
		return 1
	}
}

//对字符串数组进行排序
func insertSort(s []string, charToInt map[byte]int) {
	i, j := 0, 0
	ll := len(s)
	for i = 1; i < ll; i++ {
		tmp := s[i]
		for j := i - 1; j >= 0; j-- {
			//用给定的规则比较字符串的大小
			if compare(tmp, s[j], charToInt) == -1 {
				s[j+1] = s[j]
			} else {
				break
			}
		}
		s[j+1] = tmp
	}
}

func main() {
	s := []string{"bed", "dog", "dear", "eye"}
	sequence := "dgecfboa"
	//用来存储字母序列与其对应的值的键值对
	charToInt := map[byte]int{}
	//根据给定字符序列构造map
	for i, v := range sequence {
		charToInt[byte(v)] = i
	}
	insertSort(s, charToInt)
	fmt.Println(s)
}
