package demo17

import (
	"fmt"
)

//求解字符串中字典序最大的子序列

//顺序遍历法
//求串中字典序最大的子序列
func getLargestSub(str string) string {
	ll := len(str)
	largestSub := make([]byte, ll+1)
	k := 0
	for i := 0; i < ll; i++ {
		largestSub[k] = str[i]
		for j := i + 1; j < ll; j++ {
			//找出第i个字符后面最大的字符放到largestSub[k]中
			if str[j] > largestSub[k] {
				largestSub[k] = str[j]
				i = j
			}
		}
		k++
	}
	return string(largestSub[0:k])
}

//逆序遍历法
func getLargestSub2(txt string) string {
	ll := len(txt)
	largestSub := make([]byte, ll+1)
	//最后一个字符一定在子串中
	largestSub[0] = txt[ll-1]
	i := ll - 2
	j := 0
	//逆序遍历字符串
	for ; i >= 0; i-- {
		if txt[i] >= largestSub[j] {
			j++
			largestSub[j] = txt[i]
		}
	}
	largestSub[j+1] = ' '
	//最子串进行逆序遍历
	for i = 0; i < j; i, j = i+1, j-1 {
		tmp := largestSub[i]
		largestSub[i] = largestSub[j]
		largestSub[j] = tmp
	}
	return string(largestSub)
}

func main() {
	a := "acbdxmng"
	fmt.Println("顺序遍历法：")
	result := getLargestSub(a)
	if result == "" {
		fmt.Println("字符串为空")
	} else {
		fmt.Println(result)
	}

	fmt.Println("逆序遍历法：")
	result = getLargestSub2(a)
	if result == "" {
		fmt.Println("字符串为空")
	} else {
		fmt.Println(result)
	}

}
