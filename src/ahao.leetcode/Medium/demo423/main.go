package main

import (
	"bytes"
	"strings"
)

//423. 从英文中重建数字
func originalDigits(s string) string {
	c := map[rune]int{}
	for _, ch := range s {
		c[ch]++
	}
	cnt := [10]int{}
	cnt[0] = c['z']
	cnt[2] = c['w']
	cnt[4] = c['u']
	cnt[6] = c['x']
	cnt[8] = c['g']

	cnt[3] = c['h'] - cnt[8]
	cnt[5] = c['f'] - cnt[4]
	cnt[7] = c['s'] - cnt[6]

	cnt[1] = c['o'] - cnt[0] - cnt[2] - cnt[4]

	cnt[9] = c['i'] - cnt[5] - cnt[6] - cnt[8]

	res := []byte{}
	for i, c := range cnt {
		res = append(res, bytes.Repeat([]byte{byte('0' + i)}, c)...)
	}
	return string(res)
}

func originalDigits2(s string) string {
	digits := make([]int, 26)
	for i := 0; i < len(s); i++ {
		digits[int(s[i]-'a')]++
	}
	res := make([]string, 10)
	res[0] = convert('z', digits, "zero", "0")
	res[6] = convert('x', digits, "six", "6")
	res[2] = convert('w', digits, "two", "2")
	res[4] = convert('u', digits, "four", "4")
	res[5] = convert('f', digits, "five", "5")
	res[1] = convert('o', digits, "one", "1")
	res[7] = convert('s', digits, "seven", "7")
	res[3] = convert('r', digits, "three", "3")
	res[8] = convert('t', digits, "eight", "8")
	res[9] = convert('i', digits, "nine", "9")
	return strings.Join(res, "")
}

func convert(b byte, digits []int, s string, num string) string {
	v := digits[int(b-'a')]
	for i := 0; i < len(s); i++ {
		digits[int(s[i]-'a')] -= v
	}
	return strings.Repeat(num, v)
}
