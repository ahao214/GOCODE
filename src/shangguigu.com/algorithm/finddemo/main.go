package main


import(
	"fmt"
)

func main(){
	//定义数组
	names:=[4]string{"紫衫龙王","金毛狮王","青翼蝠王","白眉鹰王"}
	//定义变量
	var heroName=""
	fmt.Println("请输入要查找的人名：")
	fmt.Scanln(&heroName)

	//顺序查找(1)
	for i:=0;i<len(names);i++{
		if heroName==names[i]{
			fmt.Printf("找到了，名字是:%v,下标是:%v",heroName,i)
			break
		}else if i==(len(names)-1){
			fmt.Printf("没有找到:%v",heroName)
		}
	}

	//顺序查找(2)
	index := -1
	for i:=0;i<len(names);i++{
		if heroName == names[i]{			
			index = i
			break
		}
	}
	if index != -1{
		fmt.Printf("找到了，名字是:%v,下标是:%v",heroName,index)			
	}else{
		fmt.Println("没有找到:",heroName)
	}

}