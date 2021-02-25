package main

import (
	"fmt"
	"sort"
)

//题目：第881题：救生艇
//输入：people = [1,2], limit = 3
//输出：1
//解释：1 艘船载 (1, 2)
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	res := 0
	i, j := 0, len(people)-1
	for i <= j {
		if people[i]+people[j] <= limit {
			i++
		} else {
			j--
		}
		res++
	}
	return res
}

func main() {
	people := []int{1, 2, 3, 4}
	limit := 5
	fmt.Println(numRescueBoats(people, limit))
}
