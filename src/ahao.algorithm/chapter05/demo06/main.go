package main

import (
	"fmt"
)

//对由大小写字母组成的字符数组排序

func ReverseArray(ch []rune) {
	begin := 0
	end := len(ch) - 1
	for begin < end {
		//正向遍历找到下一个大写字母
		for ch[begin] >= 'a' && ch[begin] <= 'z' && end > begin {
			begin++
		}
		//逆向遍历找到下一个小写字母
		for ch[end] >= 'A' && ch[end] <= 'Z' && end > begin {
			end--
		}
		tmp := ch[begin]
		ch[begin] = ch[end]
		ch[end] = tmp
	}
}

func main() {
	ch := []rune("AbcDef")
	ReverseArray(ch)
	fmt.Println(string(ch))

}
