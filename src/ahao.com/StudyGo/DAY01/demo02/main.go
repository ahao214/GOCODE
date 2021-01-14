package main

import (
	"fmt"
)

func main() {
	//数组
	var temp [3]string
	temp[0] = "Jack"
	temp[1] = "Tom"
	temp[2] = "Henlo"
	fmt.Println(temp)

	var team1 = [3]string{"one", "two", "three"}
	var team2 = [...]int{1, 2, 3}
	fmt.Println(team1)
	fmt.Println(team2)
	for k, v := range temp {
		fmt.Println(k, v)
	}

	//切片

	//从数组生成切片
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])

	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}

	//区间
	fmt.Println(highRiseBuilding[10:15])
	//中间到末尾
	fmt.Println(highRiseBuilding[20:])
	//从开头到中间
	fmt.Println(highRiseBuilding[:15])

	//把切片的开始和结束位置设置为0，生成的切片将变成空
	c := []int{1, 2, 3}
	fmt.Println(c[0:0])

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len:%d,cap:%d\n", len(numbers), cap(numbers))
	}

	var car []string
	//添加一个元素
	car = append(car, "oldfll")
	fmt.Println(car)
	//添加多个元素
	car = append(car, "one", "two", "three")
	fmt.Println(car)
	//添加切片
	team := []string{"bus", "light"}
	car = append(car, team...)
	fmt.Println(car)

	//从切片中删除元素
	seq := []string{"a", "b", "c", "d", "e", "f", "g"}
	//删除指定位置
	index := 2
	//查看删除位置之前和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	//将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)

}
