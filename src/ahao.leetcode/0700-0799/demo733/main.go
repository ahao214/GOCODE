package main

//733.图像渲染
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	//原始颜色
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	m := len(image)
	n := 0
	if m > 0 {
		n = len(image[0])
	}
	fillColor(image, m, n, sr, sc, oldColor, newColor)
	return image
}

func fillColor(image [][]int, m, n, row, col, oldColor, newColor int) {
	if row < 0 || row >= m || col < 0 || col >= n || image[row][col] != oldColor {
		return
	}
	image[row][col] = newColor
	fillColor(image, m, n, row-1, col, oldColor, newColor)
	fillColor(image, m, n, row+1, col, oldColor, newColor)
	fillColor(image, m, n, row, col-1, oldColor, newColor)
	fillColor(image, m, n, row, col+1, oldColor, newColor)
}
