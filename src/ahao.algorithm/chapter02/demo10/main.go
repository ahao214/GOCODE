package main

import (
	"fmt"
)

type Pair struct {
	first  int
	second int
}

func FindPairs(arr []int) bool {
	//键为数对的和，值为数对
	sumPair := map[int]*Pair{}
	n := len(arr)
	//遍历数组中所有可能的数对
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			//如果这个数对的和在map中没有，则放入map中
			sum := arr[i] + arr[j]
			if _, ok := sumPair[sum]; !ok {
				sumPair[sum] = &Pair{i, j}
			} else {
				//map中已经存在于sum相同的数对了，找出来并打印出啦
				//找出已经遍历过的并存储在map中和为sum的数对
				p := sumPair[sum]
				fmt.Printf("(%v,%v),(%v,%v)", arr[p.first], arr[p.second], arr[i], arr[j])
				fmt.Println()
				return true
			}
		}
	}
	return false
}

//从数组中找出满足a+b=c+d的两个数对
func main() {
	arr := []int{3, 4, 7, 10, 20, 9, 8}
	FindPairs(arr)
}
