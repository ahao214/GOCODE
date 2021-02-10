package main

import "fmt"

func firstUniqChar(s string) int {
	var arr [26]int
	for i, k := range s {
		arr[k-'a'] = i
	}
	for i, k := range s {
		if i == arr[k-'a'] {
			return i
		} else {
			arr[k-'a'] = -1
		}
	}
	return -1
}

//387:字符串中的第一个唯一字符
func main() {
	s := "people"
	res := firstUniqChar(s)
	fmt.Println(res)
}
