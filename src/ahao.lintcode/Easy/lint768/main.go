package main

//768 · 杨辉三角
func calcYangHuisTriangle(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else if i > 1 {
				row = append(row, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, row)
	}
	return result
}
