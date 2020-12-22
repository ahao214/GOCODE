package main

import(
	"fmt"
)

//map的遍历

func main(){

	cities:=make(map[string]string)
	cities["c1"]="北京"
	cities["c2"]="上海"
	cities["c3"]="广州"
	cities["c4"]="深圳"

	for k,v:=range cities{
		fmt.Printf("k=%v,v=%v\n",k,v)
	}


	studentMap:=make(map[string]map[string]string)

	studentMap["stu01"]=make(map[string]string,3)
	studentMap["stu01"]["name"]="Tom"
	studentMap["stu01"]["sex"]="male"
	studentMap["stu01"]["address"]="北京"

	studentMap["stu02"]=make(map[string]string,3)
	studentMap["stu02"]["name"]="Andi"
	studentMap["stu02"]["sex"]="female"
	studentMap["stu02"]["address"]="上海"

	for k1,v1:=range studentMap{
		fmt.Println("k1=",k1)
		for k2,v2:=range v1{
			fmt.Printf("\t k2:%v,v2:%v",k2,v2)
		}
		fmt.Println()
	}

}