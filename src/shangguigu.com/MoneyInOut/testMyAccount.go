package MoneyInOut

import(
	"fmt"
)


func main(){
	//声明一个变量，用来接收用户的输入值
	key:=""
	//声明一个变量，控制是否退出for循环
	loop:=true
	//显示主菜单
	for{
		fmt.Println("-------家庭记账收支软件------")
		fmt.Println("		1-收支明细")
		fmt.Println("		2-登记收入")
		fmt.Println("		3-登记支出")
		fmt.Println("		4-退出软件")
		fmt.Print("请选择(1-4):")

		fmt.Scanln(&key)
		switch key{
		case "1":
			fmt.Println("---收支明细---")
		case "2":
			fmt.Println("登记收入")
		case "3":
			fmt.Println("登记支出...")
		case "4":
			loop=false
		default:
			fmt.Println("请输入正确的选项...")
		}

		if !loop{
			break
		}

	}
	fmt.Println("你已经退出该软件...")




}