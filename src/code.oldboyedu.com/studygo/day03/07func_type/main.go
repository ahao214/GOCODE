package main

import "fmt"

func f1(){
	fmt.Println("hello world")
}

func f2() int {
	return 10
}

func f3(x func() int){
	ret := x()
	fmt.Println(ret)
}

func f4(x,y,z int)int{
	return x + y + z
}

//函数类型
func main(){
	a:=f1
	fmt.Printf("%T\n",a)
	b:=f2
	fmt.Printf("%T\n",b)

	f3(f2)
	f3(b)
	

}