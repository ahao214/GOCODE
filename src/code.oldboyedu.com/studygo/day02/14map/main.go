package main

import "fmt"

func main(){

	var m1 map[string]int
	fmt.Println(m1==nil)	//还没有初始化
	m1 = make(map[string]int, 10)


	m1["Tom"]=19
	m1["Jerry"]=20

	fmt.Println(m1)

	fmt.Println(m1["Tom"])

	value,ok:=m1["Jeddry"]
	if !ok{
		fmt.Println("查无此key")
	}else{
		fmt.Println(value)
	}

	//遍历map
	for k,v:=range m1{
		fmt.Println(k,v)
	}

	//删除map里面的数据
	delete(m1,"Tom")
	fmt.Println(m1)
	

}