package main

import (
	"fmt"
	"sort"
)

//学生结构体
type StuScore struct {
	name  string
	score int
}

type StuScores []StuScore

//Len():人数
func (s StuScores) Len() int {
	return len(s)
}

//Less():成绩将由低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

//Swap():排序
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//排序
func main() {
	stus := StuScores{
		{"张三", 98},
		{"李四", 78},
		{"王五", 58},
		{"赵六", 97},
	}
	fmt.Println("---默认---")

	//原始顺序
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}

	fmt.Println()

	sort.Sort(stus)

	fmt.Println("---排序之后---")
	//排好序后的结构
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}
	//判断是否已经排好顺序，将会打印true
	fmt.Println("是否已经排序？", sort.IsSorted(stus))

}
