package main

//390. 消除游戏
func lastRemaining(n int) int {
	a1, an := 1, n
	k, cnt, step := 0, n, 1
	for cnt > 1 {
		if k%2 == 0 { // 正向
			a1 += step
			if cnt%2 == 1 {
				an -= step
			}
		} else { // 反向
			if cnt%2 == 1 {
				a1 += step
			}
			an -= step
		}
		k++
		cnt >>= 1
		step <<= 1
	}
	return a1
}
