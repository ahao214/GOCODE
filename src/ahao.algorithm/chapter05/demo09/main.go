package main

import (
	"fmt"
)

//实现字符串的匹配

//直接计算方法
func match(s, p string) int {
	slen, plen := len(s), len(p)
	//p肯定不是s的字串
	if slen < plen {
		return -1
	}
	i, j := 0, 0
	for i < slen && j < plen {
		if s[i] == p[j] {
			//如果相同，则继续比较后面的字符
			i, j = i+1, j+1
		} else {
			//后退回去重新比较
			i = i - j + 1
			j = 0
			if i > slen-plen {
				return -1
			}
		}
	}
	if j >= plen {
		//匹配成功
		return i - plen
	}
	return -1
}

func main() {
	s, p := "xyzabcd", "abc"
	fmt.Println(match(s, p))
}
