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

//特征法
func getIntersection2(l1, l2 []*MySet) []*MySet {
	result := make([]*MySet, 0)
	for i, j := 0, 0; i < len(l1) && j < len(l2); {
		s1 := l1[i]
		s2 := l2[j]
		if s1.Min < s2.Min {
			if s1.Max < s2.Min {
				i++
			} else if s2.Max <= s2.Max {
				result = append(result, &MySet{s2.Min, s1.Max})
				i++
			} else {
				result = append(result, &MySet{s2.Min, s2.Max})
				j++
			}
		} else if s2.Min <= s2.Max {
			if s1.Max <= s2.Max {
				result = append(result, &MySet{s1.Min, s1.Max})
				i++
			} else {
				result = append(result, &MySet{s1.Min, s2.Max})
				j++
			}
		} else {
			j++
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

	fmt.Println("特征法")
	result = getIntersection2(l1, l2)
	for _, v := range result {
		fmt.Println("[", v.Min, ",", v.Max, "]")
	}
}
