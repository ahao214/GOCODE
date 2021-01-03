package main


import(
	"fmt"
)


func main(){
	//声明一个变量，用来接收用户的输入值
	key:=""
	//声明一个变量，控制是否退出for循环
	loop:=true

	//账户的余额
	balance:=10000.0
	//每次收支的金额
	money:=0.0
	//每次收支的说明
	note:=""
	//定义一个变量，来记录是否有收支的行为
	flag:=false
	//收支的详情有字符串来表示
	detail:="收支\t账户金额\t收支金额\t说	明"

	//显示主菜单
	for{
		fmt.Println("\n-------家庭记账收支软件------")
		fmt.Println("		1-收支明细")
		fmt.Println("		2-登记收入")
		fmt.Println("		3-登记支出")
		fmt.Println("		4-退出软件")
		fmt.Print("请选择(1-4):")

		fmt.Scanln(&key)
		switch key{
		case "1":
			fmt.Println("---当前收支明细记录---")
			if flag{
				fmt.Println(detail)
			}else{
				fmt.Println("当前没有收支明细，来一笔吧!")
			}

		case "2":
			fmt.Println("登记收入")
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance+=money	//修改账户余额
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			//将这个收入情况拼接到detail变量
			detail+=fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note)
			flag=true
		case "3":
			fmt.Println("登记支出...")
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance{
				fmt.Println("余额不足...")
				break
			}else{
				balance-=money	//修改账户余额
			}			
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			//将这个收入情况拼接到detail变量
			detail+=fmt.Sprintf("\n支出\t%v\t%v\t%v",balance,money,note)
			flag=true
		case "4":
			fmt.Println("你确定要退出吗？y/n")
			choice:=""
			for{
				fmt.Scanln(&choice)
				if choice=="y" || choice=="n"{
					break
				}
				fmt.Println("你的输入有误，请重新输入! y/n")
			}
			if choice=="y"{
				loop=false
			}
			
		default:
			fmt.Println("请输入正确的选项...")
		}

		if !loop{
			break
		}

	}
	fmt.Println("你已经退出该软件...")




}