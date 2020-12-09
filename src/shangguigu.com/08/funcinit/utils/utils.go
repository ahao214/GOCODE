package utils

import(
	"fmt"
)

var Age int
var Name string

//init函数完成初始化工作
func init(){
	fmt.Println("utils包里面的init()函数")
	Age=100
	Name="Tom"
}