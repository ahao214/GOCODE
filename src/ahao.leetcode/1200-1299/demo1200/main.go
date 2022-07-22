package main

/*
1200. 最小绝对差
*
func minimumAbsDifference(arr []int) (ans [][]int) {
	sort.Ints(arr)
	for i, best := 0, math.MaxInt32; i < len(arr)-1; i++ {
		if delta := arr[i+1] - arr[i]; delta < best {
			best = delta
			ans = [][]int{{arr[i], arr[i+1]}}
		} else if delta == best {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}
	return
}
