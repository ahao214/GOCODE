package main

import (
	"fmt"
)

//切片语法返回子串
func main() {
	s := "abcdefg"
	one := s[:3]
	two := s[1:4]
	three := s[2:]
	fmt.Println(one, two, three)

	s = "枯藤"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d:[%c]\n", i, s[i])
	}

	for i, c := range s {
		fmt.Printf("%d:[%c]\n", i, c)
	}

}
