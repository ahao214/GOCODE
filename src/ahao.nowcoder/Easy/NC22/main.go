package main

/**
 * NC22 合并两个有序的数组
 * @param A int整型一维数组
 * @param B int整型一维数组
 * @return void
 */
func merge(A []int, m int, B []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if A[m-1] >= B[n-1] {
			A[p-1] = A[m-1]
			m--
		} else {
			A[p-1] = B[n-1]
			n--
		}
	}
	for ; n > 0; n-- {
		A[n-1] = B[n-1]
	}
}
