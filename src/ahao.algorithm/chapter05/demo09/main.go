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

//KMP算法
func getNext(p string, next []int) {
	next[0] = -1
	for i, j := 0, -1; i < len(p); {
		if j == -1 || p[i] == p[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
}

func matchKMP(s, p string, next []int) int {
	slen, plen := len(s), len(p)
	//p肯定不是s的字串
	if slen < plen {
		return -1
	}
	i, j := 0, 0
	for i < slen && j < plen {
		fmt.Printf("i=%v,j=%v", i, j)
		fmt.Println()
		if j == -1 || s[i] == p[j] {
			//如果相同，则继续比较后面的字符
			i++
			j++
		} else {
			//主串i不需要回溯，从next数组中找出需要比较的模式串的位置j
			j = next[j]
		}
	}
	if j >= plen {
		return i - plen
	}
	return -1
}

func main() {
	s, p := "xyzabcd", "abc"
	fmt.Println(match(s, p))

	fmt.Println("KMP")
	s, p = "abababaabcbab", "abaabc"
	next := make([]int, len(p)+1)
	getNext(p, next)
	fmt.Println("next数组为；", next)
	fmt.Println("匹配结果为：", matchKMP(s, p, next))
}
