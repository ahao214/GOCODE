package main

//2022. 将一维数组转变成二维数组
func construct2DArray(original []int, m, n int) [][]int {
	k := len(original)
	if k != m*n {
		return nil
	}
	ans := make([][]int, 0, m)
	for i := 0; i < k; i += n {
		ans = append(ans, original[i:i+n])
	}
	return ans
}

func construct2DArray2(original []int, m, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	res := make([][]int, m)

	left, right := 0, n
	for i := 0; i < m; i++ {
		res[i] = original[left:right]
		left += n
		right += n
	}
	return res
}
