package main

import (
	"fmt"
	"strings"
)

func main() {
	m := map[string]int{
		"abc": 123,
	}

	key := []byte("abc")
	x, ok := m[string(key)]
	fmt.Println(x, ok)
}

func Join(a []string, sep string) string {
	//统计分隔符长度
	n := len(sep) * (len(a) - 1)
	//统计所有待拼接字符串长度
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}
	//一次分配所需长度的数组空间
	b := make([]byte, n)
	//拷贝数据
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func test() string {
	s := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return strings.Join(s, "")
}
