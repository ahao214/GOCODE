package main


import(
	"fmt"
)

//定义一个结构体
type Cat struct{
	Name string
	Age int
	Color string
}

func main(){
	var cat1 Cat
	cat1.Name="小白"
	cat1.Age=12
	cat1.Color="red"
	fmt.Println(cat1)

	fmt.Printf("cat1的地址是%p\n",&cat1)

	fmt.Println("猫的信息如下：")
	fmt.Println("猫的名字是：",cat1.Name)
	fmt.Println("猫的年龄是：",cat1.Age)
	fmt.Println("猫的颜色是：",cat1.Color)

}

