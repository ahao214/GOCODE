package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	//解析指定文件生成模板对象
	tmp, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("create template failed,err:", err)
		return
	}
	//利用给定数据渲染模板，并将结果写入
	user := UserInfo{
		Name:   "枯藤",
		Gender: "男",
		Age:    30,
	}
	tmp.Execute(w, user)
}

func main() {

}
