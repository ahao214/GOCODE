package main

import (
	"fmt"
)

type Map map[string]string

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type iMap Map

//只要底层类型是slice/map等支持range的类型字面量

func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type slice []int

func (s slice) Print() {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "tata"
	//mp与ma有相同的底层类型map[string]string,并mp是未命名类型
	//所以mp可以直接赋值给ma
	var ma Map = mp

	ma.Print()

	var i interface {
		Print()
	} = ma
	i.Print()

	s1 := []int{1, 2, 3}
	var s2 slice
	s2 = s1
	s2.Print()
}
