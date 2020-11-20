package main


import "fmt"


func deferDemo(){
	fmt.Println("start")
	defer fmt.Println("hello")	//有多个defer，先进后出
	defer fmt.Println("world")
	fmt.Println("end")


}

func main(){
	deferDemo()



}