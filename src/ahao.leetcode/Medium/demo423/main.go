package main

import "bytes"

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
