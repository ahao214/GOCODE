package main

import(
	"fmt"
)


func TypeJudge(items... interface{}){
	for index,x:=range items{
		index+=1
		switch x.(type){
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v\n",index,x)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n",index,x)
		case float64:
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n",index,x)
		case int,int32,int64:
			fmt.Printf("第%v个参数是 整数 类型，值是%v\n",index,x)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v\n",index,x)	
		default:
			fmt.Printf("第%d个参数是类型不确定，值是%v\n",index,x)	

		}
	}
}

func main(){
	var n1 float32=1.1
	var n2 float64=2.3
	var n3 int32=30
	var name="JACK"

	address:="深圳"
	n4:=400
	TypeJudge(n1,n2,n3,name,address,n4)


}