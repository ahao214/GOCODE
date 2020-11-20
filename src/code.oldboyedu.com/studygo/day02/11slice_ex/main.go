package main

import (
	"fmt"
	"sort"

)

func main(){
	var a = make([]int,5, 10)
	fmt.Println(a)
	for i:=0;i<10;i++{
		a = append(a,i)
	}
	fmt.Println(a)

	s := [...]int{1,32,4,7,5}
	sort.Ints(s[:])	//对切片排序
	fmt.Println(s)

}