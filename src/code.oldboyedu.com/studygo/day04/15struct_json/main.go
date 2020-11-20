package main

import (
	"fmt"
	"encoding/json"

)

//结构体和json
type person struct{
	Name string `json:"name"`
	Age int		`json:"age"`	
}


func main(){
	p:=person{
		Name:"Tom",
		Age:20,
	}
	//序列化
	b,err:=json.Marshal(p)
	if err!=nil{
		fmt.Printf("marshal failed, err:%v",err)
		return
	}
	fmt.Printf("%#v\n",string(b))

	//反序列化
	str:=`{"name":"Jerry","age":19}`
	var p1 person
	json.Unmarshal([]byte(str),&p1)	//传指针是为了能修改p1的值
	fmt.Printf("%#v\n",p1)
}