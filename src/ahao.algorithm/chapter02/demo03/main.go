package main

import(
	"fmt"
	"os"
)

var (
	HOME=os.Getenv("HOME")
	USER=os.Getenv("USER")
	GOROOT=os.Getenv("GOROOT")
)

func GetClass()(stuNum int,className string,headTeacher string){
	return 40,"一班","李四"
}

func main(){
	fmt.Println(HOME,USER,GOROOT)

	stuNum,_,_:=GetClass()
	fmt.Println(stuNum)

}