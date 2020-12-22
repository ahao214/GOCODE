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

	//查询
	val,ok:=cities["c1"]
	if ok{
		fmt.Printf("c1的值是%v\n",val)
	}else{
		fmt.Println("c1没有值")
	}

	//delete删除的时候，key存在，直接删除，key不存在，删除不会操作，也不会报错
	delete(cities,"c1")
	delete(cities,"c4")
	fmt.Println(cities)

	//删除所有的值
	//1.变量map，删除值
	//2.直接make一个新的空间，让原来的成为垃圾，被gc回收
	cities=make(map[string]string)
	fmt.Println(cities)



}