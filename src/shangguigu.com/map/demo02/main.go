package main

import(
	"fmt"
)

func main(){
	studentMap:=make(map[string]map[string]string)

	studentMap["stu01"]=make(map[string]string,3)
	studentMap["stu01"]["name"]="Tom"
	studentMap["stu01"]["sex"]="male"
	studentMap["stu01"]["address"]="北京"

	studentMap["stu02"]=make(map[string]string,3)
	studentMap["stu02"]["name"]="Andi"
	studentMap["stu02"]["sex"]="female"
	studentMap["stu02"]["address"]="上海"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu01"])
	fmt.Println(studentMap["stu02"])
	fmt.Println(studentMap["stu02"]["name"])
}