package main

/*
1282. 用户分组
*/
func groupThePeople(groupSizes []int) (ans [][]int) {
	groups := map[int][]int{}
	for i, size := range groupSizes {
		groups[size] = append(groups[size], i)
	}
	for size, people := range groups {
		for i := 0; i < len(people); i += size {
			ans = append(ans, people[i:i+size])
		}
	}
	return
}
