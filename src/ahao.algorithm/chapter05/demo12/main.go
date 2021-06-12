package main

import (
	"fmt"
)

//判断一个字符串是否包含重复字符

//蛮力法

func isDup(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return true
			}
		}
	}
	return false
}

//空间换时间
func isDup2(str string) bool {
	//只需要8个32位的int，8*32=256位
	flag := make([]int, 8)
	for _, v := range str {
		index := int(v) / 32
		shift := uint(v) % 32
		if (flag[index] & (1 << shift)) != 0 {
			return true
		}
		flag[index] |= 1 << shift
	}
	return false
}

func main() {
	str := "GODO"
	fmt.Println("方法一：")
	if isDup(str) {
		fmt.Println(str, "中有重复字符")
	} else {
		fmt.Println(str, "中没有重复字符")
	}

	fmt.Println("方法二：")
	if isDup2(str) {
		fmt.Println(str, "中有重复字符")
	} else {
		fmt.Println(str, "中没有重复字符")
	}

}
