package main

import (
	"fmt"
)

//& 取出地址
//* 根据地址取出地址指向的值

func main() {
	//指针取值
	a := 10
	b := &a //取变量a的地址，将指针保存到b中
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p,type:%T\n", b, b)
	fmt.Println(&b)

	c := *b //指针取值(根据指针去内存取值)
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	s := 10
	modifyone(s)
	fmt.Println(s)
	modifytwo(&s)
	fmt.Println(s)

	//空指针
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是：%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}

}

//指针传值示例
func modifyone(x int) {
	x = 100
}

func modifytwo(x *int) {
	*x = 100
}
