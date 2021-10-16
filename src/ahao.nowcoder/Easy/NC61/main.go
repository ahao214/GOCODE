package main

/**
 * 两数之和
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		another := target - numbers[i]
		if _, ok := m[another]; ok {
			if i < m[another] {
				return []int{i, m[another]}
			} else {
				return []int{m[another], i}
			}
		}
		m[numbers[i]] = i
	}
	return nil
}
