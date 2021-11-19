package main

//NC175 整数反转
func reverse(num int) int {
	tmp := 0
	for num != 0 {
		tmp = tmp*10 + num%10
		num = num / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}
