package main

import(
	"fmt"
	"shangguigu.com/FamilyAccount/utils"
)

func main(){
	fmt.Println("这是面向对象方式完成的！")
	utils.NewFamilyAccount().MainMenu()
}