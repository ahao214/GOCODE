package main

import(
	"fmt"
)

//map切片
func main(){

	//声明一个map切片
	var monsters [] map[string]string
	monsters=make([]map[string]string,2)
	//增加一个信息
	if monsters[0]==nil{
		monsters[0]=make(map[string]string,2)
		monsters[0]["name"]="牛魔王"
		monsters[0]["age"]="500"
	}

	if monsters[1]==nil{
		monsters[1]=make(map[string]string,2)
		monsters[1]["name"]="玉兔精"
		monsters[1]["age"]="400"
	}	

	newMonster:=map[string]string{
		"name":"火云邪神",
		"age":"250",
	}
	//使用append动态增加
	monsters=append(monsters,newMonster)

	fmt.Println(monsters)
	
}