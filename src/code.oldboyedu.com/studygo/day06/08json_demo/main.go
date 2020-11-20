package main

import(
	"fmt"
	"encoding/json"
)

type person struct{
	Name string	`json:"name"`
	Age int		`json:"age"`

}

func main(){

	str:=`{"name":"ahao","age":90}`
	var p person
	json.Unmarshal([]byte(str),&p)
	fmt.Println(p.Name,p.Age)
}