package main

//240. 搜索二维矩阵 II
func searchMatrixOne(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for col >= 0 && row <= len(matrix)-1 {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			row++
		} else {
			col--
		}
	}
	return false
}

func searchMatrixTwo(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for _, row := range matrix {
		low, high := 0, len(matrix[0])-1
		for low <= high {
			mid := low + (high-low)>>1
			if row[mid] > target {
				high = mid - 1
			} else if row[mid] < target {
				low = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}
