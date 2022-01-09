package main

//1629. 按键持续时间最长的键
func slowestKey(releaseTimes []int, keysPressed string) byte {
	ans := keysPressed[0]
	longestTime := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		time := releaseTimes[i] - releaseTimes[i-1]
		key := keysPressed[i]
		if time > longestTime || time == longestTime && key > ans {
			ans = key
			longestTime = time
		}
	}
	return ans
}
