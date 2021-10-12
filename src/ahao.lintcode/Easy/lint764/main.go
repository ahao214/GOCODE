package main

/**
 * 764 · 计算圆周长和面积
 * @param r: a Integer represent radius
 * @return: the circle's circumference nums[0] and area nums[1]
 */
func calculate(r int) []float64 {
	res := []float64{}
	const pi = 3.14
	perimeter := float64(pi) * 2 * float64(r)
	area := float64(pi) * float64(r) * float64(r)
	res = append(res, perimeter)
	res = append(res, area)
	return res
}
