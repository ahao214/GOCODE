package main

/*
1345. 跳跃游戏 IV
给你一个整数数组 arr ，你一开始在数组的第一个元素处（下标为 0）。

每一步，你可以从下标 i 跳到下标：

i + 1 满足：i + 1 < arr.length
i - 1 满足：i - 1 >= 0
j 满足：arr[i] == arr[j] 且 i != j
请你返回到达数组最后一个元素的下标处所需的 最少操作次数 。

注意：任何时候你都不能跳到数组外面。
*/
func minJumps(arr []int) int {
	n := len(arr)
	idx := map[int][]int{}
	for i, v := range arr {
		idx[v] = append(idx[v], i)
	}
	vis := map[int]bool{0: true}
	type pair struct{ idx, step int }
	q := []pair{{}}
	for {
		p := q[0]
		q = q[1:]
		i, step := p.idx, p.step
		if i == n-1 {
			return step
		}
		for _, j := range idx[arr[i]] {
			if !vis[j] {
				vis[j] = true
				q = append(q, pair{j, step + 1})
			}
		}
		delete(idx, arr[i])
		if !vis[i+1] {
			vis[i+1] = true
			q = append(q, pair{i + 1, step + 1})
		}
		if i > 0 && !vis[i-1] {
			vis[i-1] = true
			q = append(q, pair{i - 1, step + 1})
		}
	}
}
