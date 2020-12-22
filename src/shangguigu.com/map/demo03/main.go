package main

import(
	"fmt"
)

//map的增删改查

func main(){
	
	cities:=make(map[string]string)
	cities["c1"]="上海"
	cities["c2"]="北京"
	cities["c3"]="广州"
	//key值存在就是修改，不存在，就是新增
	cities["c3"]="深圳"

	fmt.Println(cities)
}