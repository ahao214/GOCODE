package main

import (
	"fmt"

	"code.oldboyedu.com/studygo/day05/calc"

)

var x =10
const pi =3.14
func init(){
	fmt.Println("自动执行")
	fmt.Println(x)
	fmt.Println(pi)
}

func main(){
	ret:=calc.Add(10,20)
	fmt.Println(ret)

}