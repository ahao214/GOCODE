package main

import (
	"fmt"

)

var valueone float64
func main(){

	s:="abce你好"
	for i:=0;i<len(s);i++{
		fmt.Println(s[i])
	}
	c:=s[len(s)]
	fmt.Println(len(s),c)
	fmt.Println(s[0:7])


}