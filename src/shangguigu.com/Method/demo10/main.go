package main

import(
	"fmt"
)

type Visitor struct{
	Name string
	Age int
}

func(visitor *Visitor)showPrice(){
	if visitor.Age>=90 || visitor.Age<=8{
		fmt.Println("安全第一，不要游玩")		
	}else if(visitor.Age>18){
		fmt.Printf("游客的名字是%v,年龄是%v,收费20元\n",visitor.Name,visitor.Age)
	}else{
		fmt.Printf("游客的名字是%v,年龄是%v,免费\n",visitor.Name,visitor.Age)
	}
}

func main(){

	var visitor Visitor
	for{
		fmt.Println("请输入你的名字")
		fmt.Scanln(&visitor.Name)
		if visitor.Name=="n"{
			fmt.Println("退出程序...")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&visitor.Age)
		visitor.showPrice()
	}

}