package main

import(
	"fmt"
)

func modifyUse(users map[string]map[string]string,name string){
	//判断users中是否有name
	if users[name]!=nil{
		//有这个用户,将用户的密码改为888
		users[name]["pwd"]="888"
	}else{
		//没有这个用户
		users[name]=make(map[string]string)
		users[name]["pwd"]="888"
		users[name]["nickname"]="昵称~"+name
	}
}


func main(){
	users:=make(map[string]map[string]string,10)
	users["Jack"]=make(map[string]string)
	users["Jack"]["pwd"]="999"
	users["Jack"]["nickname"]="小花猫"

	modifyUse(users,"Jack")
	modifyUse(users,"Tom")
	modifyUse(users,"Jerry")

	fmt.Println(users)

}