package main

import(
	"fmt"
	"shangguigu.com/CustomerManage/service"
)

type customerView struct{
	//定义必要字段
	key string 	//接收用户的输入
	loop bool 	//表示是否循环的显示主菜单
	customerService *service.CustomerService
}


//显示所有的客户信息
func(this *customerView)list(){
	//首先，获取到当前所有客户信息（在切片中）
	customers:=this.customerService.List()
	//显示
	fmt.Println("------客户列表------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	for i:=0;i<len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n------客户列表完成------\n\n")
}

//得到用户的输入信息，构建新的客户，并完成添加
func(this *customerView)add(){
	fmt.Println("------添加客户------")



	fmt.Println("------添加客户完成------")
}



//显示主菜单
func(this *customerView)mainMenu(){
	for{
		fmt.Println("------客户信息管理软件------")
		fmt.Println("			1 添加客户")
		fmt.Println("			2 修改客户")
		fmt.Println("			3 删除客户")
		fmt.Println("			4 客户列表")
		fmt.Println("			5 退    出")
		fmt.Print("请选择(1-5):")

		fmt.Scanln(&this.key)
		switch this.key{
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
		case "5":
			this.loop=false
		default:
			fmt.Println("输入错误，请重新输入")
		}	

		if !this.loop{
			break
		}	
	}
	fmt.Println("你退出了客户信息管理系统！")
}


func main(){
	//创建一个customerView
	customerView:=customerView{
		key:"",
		loop:true,
	}

	//完成对customerView结构体的customerService字段的初始化
	customerView.customerService=service.NewCustomerService()

	//显示主菜单
	customerView.mainMenu()
}