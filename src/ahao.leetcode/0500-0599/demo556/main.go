package main

/*
556. 下一个更大元素 III
*/

//方法一：下一个排列
func nextGreaterElement(n int) int {
	nums := []byte(strconv.Itoa(n))
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}

	j := len(nums) - 1
	for j >= 0 && nums[i] >= nums[j] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
	ans, _ := strconv.Atoi(string(nums))
	if ans > math.MaxInt32 {
		return -1
	}
	return ans
}

func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}


//方法二：数学
func nextGreaterElement(n int) int {
	x, cnt := n, 1
	for ; x >= 10 && x/10%10 >= x%10; x /= 10 {
		cnt++
	}
	x /= 10
	if x == 0 {
		return -1
	}

	targetDigit := x % 10
	x2, cnt2 := n, 0
	for ; x2%10 <= targetDigit; x2 /= 10 {
		cnt2++
	}
	x += x2%10 - targetDigit // 把 x2%10 换到 targetDigit 上

	for i := 0; i < cnt; i++ { // 反转 n 末尾的 cnt 个数字拼到 x 后
		d := targetDigit
		if i != cnt2 {
			d = n % 10
		}
		if x > math.MaxInt32/10 || x == math.MaxInt32/10 && d > 7 {
			return -1
		}
		x = x*10 + d
		n /= 10
	}
	return x
}

