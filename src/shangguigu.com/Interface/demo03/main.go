package main

import(
	"fmt"
	"sort"
	"math/rand"
)

//1.声明Hero结构体
type Hero struct{
	Name string
	Age int
}

//2.声明一个Hero结构体的切片类型
type HeroSlice []Hero

//3.实现Interface接口
func(hs HeroSlice)Len()int{
	return len(hs)
}

//按照年龄Age升序排列
func(hs HeroSlice)Less(i,j int)bool{
	return hs[i].Age<hs[j].Age
}

func(hs HeroSlice)Swap(i,j int){
	// temp:=hs[i]
	// hs[i]=hs[j]
	// hs[j]=temp
	//下面这句话等价于上面的三句话
	hs[i],hs[j]=hs[j],hs[i]
}

func main(){
	var intSlice=[]int{0,-1,7,10,90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对结构体切片进行排序
	var heros HeroSlice
	for i:=0;i<10;i++{
		hero:=Hero{
			Name:fmt.Sprintf("英雄~%d",rand.Intn(100)),
			Age:rand.Intn(100),
		}
		heros=append(heros,hero)
	}
	//排序前的
	for _,v:=range heros{
		fmt.Println("排序前：",v)
	}
	fmt.Println("排序后的效果")
	//排序后的
	sort.Sort(heros)
	for _,v:=range heros{
		fmt.Println("排序后：",v)
	}
}