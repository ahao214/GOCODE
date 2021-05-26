package main

import (
	"fmt"
)

//求两个有序集合的交集

//蛮力法

type MySet struct {
	Min int
	Max int
}

func getIntersectionSub(s1, s2 *MySet) *MySet {
	if s1.Min < s2.Min {
		if s1.Max < s2.Min {
			return nil
		} else if s1.Max <= s2.Max {
			return &MySet{s2.Min, s1.Max}
		} else {
			return &MySet{s2.Min, s2.Max}
		}
	} else if s1.Min < s2.Max {
		if s1.Max <= s2.Max {
			return &MySet{s1.Min, s1.Max}
		} else {
			return &MySet{s1.Min, s2.Max}
		}
	}
	return nil
}

func getIntersection(l1, l2 []*MySet) []*MySet {
	result := make([]*MySet, 0)
	for _, v1 := range l1 {
		for _, v2 := range l2 {
			s := getIntersectionSub(v1, v2)
			if s != nil {
				result = append(result, s)
			}
		}
	}
	return result
}

func main() {
	l1 := []*MySet{
		&MySet{4, 8},
		&MySet{9, 13},
	}
	l2 := []*MySet{
		&MySet{6, 12},
	}
	fmt.Println("暴力法")
	result := getIntersection(l1, l2)
	for _, v := range result {
		fmt.Println("[", v.Min, ",", v.Max, "]")
	}
}
