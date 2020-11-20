package main

import "fmt"

//空接口作为函数参数
func show(a interface{}){
	fmt.Printf("type:%T,value:%v\n",a,a)
}

func main(){
	var m map[string]interface{}
	m = make(map[string]interface{},16)

	m["name"]="Jerry"
	m["age"]=16
	m["married"]=true
	m["hobby"]=[...]string{"run","bike","sing"}
	fmt.Println(m)

	show(false)
	show(nil)
	show(m)

}