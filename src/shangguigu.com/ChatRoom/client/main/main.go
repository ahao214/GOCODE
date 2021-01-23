package main

import (
	"fmt"
	"os"
	"shangguigu.com/ChatRoom/client/process"
)

//定义两个全局变量，一个表示用户的ID，一个表示用户的密码
var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	//var loop = true

	for true {
		fmt.Println("------欢迎登录多人聊天系统------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1~3)")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的ID：")
			fmt.Scanf("%d\n", &userId)

			fmt.Println("请输入用户的密码；")
			fmt.Scanf("%s\n", userPwd)
			//loop = false
			//完成登录
			//1.创建一个UserProcess实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户姓名(nickname)")
			fmt.Scanf("%s\n", &userName)
			//调用UserProcess，完成注册
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("退出系统")
			//loop=false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	//根据用户的输入，显示新的提示信息
	//if key == 1 {
	//	//说明用户要登录
	//	fmt.Println("请输入用户的ID：")
	//	fmt.Scanf("%d\n", &userId)
	//
	//	fmt.Println("请输入用户的密码；")
	//	fmt.Scanf("%s\n", userPwd)
	//	//因为使用了新的程序结构,
	//
	//	//这里我们需要重新调用
	//	//main2.login(userId, userPwd)
	//	//if err != nil {
	//	//	fmt.Println("登录失败")
	//	//} else {
	//	//	fmt.Println("登录成功")
	//	//}
	//} else if key == 2 {
	//	fmt.Println("进行用户注册的代码")
	//}

}
