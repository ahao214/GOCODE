package main

/*
768. 最多能完成排序的块 II
*/

//排序+哈希表
func maxChunksToSorted(arr []int) (ans int) {
	cnt := map[int]int{}
	b := append([]int{}, arr...)
	sort.Ints(b)
	for i, x := range arr {
		cnt[x]++
		if cnt[x] == 0 {
			delete(cnt, x)
		}
		y := b[i]
		cnt[y]--
		if cnt[y] == 0 {
			delete(cnt, y)
		}
		if len(cnt) == 0 {
			ans++
		}
	}
	return
}


//单调栈
func maxChunksToSorted2(arr []int) int {
	st := []int{}
	for _, x := range arr {
		if len(st) == 0 || x >= st[len(st)-1] {
			st = append(st, x)
		} else {
			mx := st[len(st)-1]
			st = st[:len(st)-1]
			for len(st) > 0 && st[len(st)-1] > x {
				st = st[:len(st)-1]
			}
			st = append(st, mx)
		}
	}
	return len(st)
}
