package main

/**
 * 214.数字的最大值
 * @param A: An integer
 * @return: a float number
 */
func maxOfArray(A []float32) float32 {
	var maxValue float32 = A[0]
	for i := 1; i < len(A); i++ {
		if A[i] > maxValue {
			maxValue = A[i]
		}
	}
	return maxValue
}
