package main

//495. 提莫攻击
func findPoisonedDuration(timeSeries []int, duration int) int {
	if timeSeries == nil || len(timeSeries) == 0 || duration == 0 {
		return 0
	}
	res := 0
	expired := 0
	for _, time := range timeSeries {
		if time >= expired {
			res += duration
		} else {
			res += time + duration - expired
		}
		expired = duration + time
	}
	return res
}
