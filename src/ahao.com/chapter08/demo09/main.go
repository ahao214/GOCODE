package main

import (
	"fmt"
)

type firstS struct {
	in1 int
	in2 int
}

type secondS struct {
	b      int
	c      float32
	int    //匿名字段
	firstS //匿名字段
}

func main() {
	sec := new(secondS)
	sec.b = 6
	sec.c = 7.4
	sec.int = 60
	sec.in1 = 1
	sec.in2 = 2
	fmt.Printf("sec.b is:%d\n", sec.b)
	fmt.Printf("sec.c is:%f\n", sec.c)
	fmt.Printf("sec.int is:%d\n", sec.int)
	fmt.Printf("sec.in1 is:%d\n", sec.in1)
	fmt.Printf("sec.in2 is:%d\n", sec.in2)

	//使用结构体字面量
	sec2 := secondS{6, 7.6, 50, firstS{5, 10}}
	fmt.Println("sec2 is:", sec2)

}
