package main

import (
	"fmt"
)

func truncateStr(str string, n int) string {
	arr := []rune(str)
	return string(arr[:n])
}

//截取包含中文的字符串
func main() {
	str := "人 ABC 们 DEF"
	fmt.Println(truncateStr(str, 6))
}
