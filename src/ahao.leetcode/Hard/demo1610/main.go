package main

import (
	"math"
	"sort"
)

//1610. 可见点的最大数目

//二分查找
func visiblePoints(points [][]int, angle int, location []int) int {
	sameCnt := 0
	polarDegrees := []float64{}
	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			sameCnt++
		} else {
			polarDegrees = append(polarDegrees, math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0])))
		}
	}
	sort.Float64s(polarDegrees)

	n := len(polarDegrees)
	for _, deg := range polarDegrees {
		polarDegrees = append(polarDegrees, deg+2*math.Pi)
	}

	maxCnt := 0
	degree := float64(angle) * math.Pi / 180
	for i, deg := range polarDegrees[:n] {
		j := sort.Search(n*2, func(j int) bool { return polarDegrees[j] > deg+degree })
		if j-i > maxCnt {
			maxCnt = j - i
		}
	}
	return sameCnt + maxCnt
}

//滑动窗口
func visiblePoints1(points [][]int, angle int, location []int) int {
	sameCnt := 0
	polarDegrees := []float64{}
	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			sameCnt++
		} else {
			polarDegrees = append(polarDegrees, math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0])))
		}
	}
	sort.Float64s(polarDegrees)

	n := len(polarDegrees)
	for _, deg := range polarDegrees {
		polarDegrees = append(polarDegrees, deg+2*math.Pi)
	}

	maxCnt := 0
	right := 0
	degree := float64(angle) * math.Pi / 180
	for i, deg := range polarDegrees[:n] {
		for right < n*2 && polarDegrees[right] <= deg+degree {
			right++
		}
		if right-i > maxCnt {
			maxCnt = right - i
		}
	}
	return sameCnt + maxCnt
}

func visiblePoints2(points [][]int, angle int, location []int) int {
	same := 0
	//转换
	angles := []float64{}
	a := 0.0
	//计算在一条直线上的点
	for _, point := range points {
		if point[0] == location[0] && point[1] == location[1] {
			same++
			continue
		}
		//atan2 可以直接鉴别象限
		a = math.Atan2(float64(point[1]-location[1]), float64(point[0]-location[0])) * 180 / math.Pi
		//转换为360度表现形式
		if a < 0 {
			a = a + 360
		}
		angles = append(angles, a)
	}
	//排序
	sort.Float64s(angles)
	length := len(angles)
	//在加一圈
	for _, angle := range angles {
		angles = append(angles, angle+360)
	}
	max := 0
	//双指针
	for l, r := 0, 0; l < length; l++ {
		for r+1 < len(angles) && (angles[r+1]-angles[l]) <= float64(angle)+1e-8 {
			r++
		}
		max = int(math.Max(float64(max), float64(r-l+1)))
	}
	return max + same
}
