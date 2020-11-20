package main

import(
	"fmt"
	"reflect"
)

type Cat struct{

}

func reflectType(x interface{}){
	v:=reflect.TypeOf(x)
	fmt.Printf("type:%v\n",v)
	fmt.Printf("type name:%v type kind:%v\n",v.Name(),v.Kind())
}

func reflectValue1(x interface{}){
	v:=reflect.ValueOf(x)
	if v.Kind()==reflect.Int64{
		v.SetInt(200)
	}
}



func reflectValue2(x interface{}){
	v:=reflect.ValueOf(x)
	if v.Elem().Kind()==reflect.Int64{
		v.Elem().SetInt(200)
	}
}
func main(){
	var a float32=3.14
	reflectType(a)
	var b int64=100
	reflectType(b)

	var c= Cat{}
	reflectType(c)

	reflectValue2(&b)
	fmt.Println(b)

}