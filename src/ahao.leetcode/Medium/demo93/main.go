package main

import (
	"strconv"
	"strings"
)

/*
93. 复原 IP 地址

有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
*/

func restoreIpAddresses(s string) []string {
	res := []string{}
	if len(s) < 4 || len(s) > 16 {
		return res
	}
	index := 0
	cur := []string{}
	backtrack(s, index, cur, &res)
	return res
}

func backtrack(s string, index int, cur []string, res *[]string) {
	if len(cur) == 4 && index == len(s) {
		*res = append(*res, strings.Join(cur, "."))
		return
	}
	for i := 1; i < 4; i++ {
		if index+i > len(s) {
			break
		}
		tmp := s[index : index+i]
		if isValidIP(tmp, i) {
			//add
			cur = append(cur, tmp)
			//backtrack
			backtrack(s, index+i, cur, res)
			//remove
			cur = cur[:len(cur)-1]
		}
	}

}

func isValidIP(ip string, index int) bool {
	if len(ip) == 0 {
		return false
	}
	if ip[0] == '0' && len(ip) > 1 {
		return false
	}
	value, err := strconv.Atoi(ip)
	if err != nil {
		return false
	}
	if value > 255 && index == 3 {
		return false
	}
	return true
}
