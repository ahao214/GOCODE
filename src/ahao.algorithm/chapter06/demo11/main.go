package main

import (
	"fmt"
	"strconv"
)

//处理数字相加的进位
func carry(num []rune, pos int) {
	for ; pos > 0; pos-- {
		if num[pos] > '9' {
			num[pos] = '0'
			num[pos-1]++
		}
	}
}

//从右到左的贪心法
//获取大于n的最小不重复数
func findMinNonDupNum1(n int) int {
	count := 0 //用来记录循环次数
	mRune := []rune(strconv.Itoa(n + 1))
	ch := make([]rune, len(mRune)+2)
	ch[0] = '0'
	ch[len(ch)-1] = '0'
	for i := 0; i < len(mRune); i++ {
		ch[i+1] = mRune[i]
	}
	ll := len(ch)
	i := ll - 2 //从右向左遍历
	for i > 0 {
		count++
		if ch[i-1] == ch[i] {
			ch[i]++      //末尾数字加1
			carry(ch, i) //处理进位
			//把下标为i后面的字符串变为0101...串
			for j := i + 1; j < ll; j++ {
				if (j-i)%2 == 1 {
					ch[j] = '0'
				} else {
					ch[j] = '1'
				}
			}
			//第i位加1后，可能会有第i+1位相等，i++用来处理这种情况
			i++
		} else {
			i--
		}
	}
	fmt.Println("循环次数为：", count)
	result, _ := strconv.Atoi(string(ch))
	return result
}

//从左到右的贪心法
func findMinNonDupNum2(n int) int {
	count := 0
	mRune := []rune(strconv.Itoa(n + 1))
	ch := make([]rune, len(mRune)+2)
	ch[0] = '0'
	ch[len(ch)-1] = '0'
	for i := 0; i < len(mRune); i++ {
		ch[i+1] = mRune[i]
	}
	ll := len(ch)
	i := 2 //从左向右遍历
	for i < ll {
		count++
		if ch[i-1] == ch[i] {
			ch[i]++      //末尾数字加1
			carry(ch, i) //处理进位
			//把下标为i后面的字符串变为0101...串
			for j := i + 1; j < ll; j++ {
				if (j-i)%2 == 1 {
					ch[j] = '0'
				} else {
					ch[j] = '1'
				}
			}
		} else {
			i++
		}
	}
	fmt.Println("循环次数为：", count)
	result, _ := strconv.Atoi(string(ch))
	return result
}

//找最小的不重复数
func main() {
	fmt.Println("从右到左的贪心法")
	fmt.Println(findMinNonDupNum1(23345))
	fmt.Println(findMinNonDupNum1(1101010))
	fmt.Println(findMinNonDupNum1(99010))
	fmt.Println(findMinNonDupNum1(8989))

	fmt.Println("从左到右的贪心法")
	fmt.Println(findMinNonDupNum2(23345))
	fmt.Println(findMinNonDupNum2(1101010))
	fmt.Println(findMinNonDupNum2(99010))
	fmt.Println(findMinNonDupNum2(8989))

}
