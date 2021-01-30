package main

import (
	"fmt"
)

//小孩的结构体
type Boy struct {
	No   int  //编号
	Next *Boy //指向下一个小孩的指针
}

//编写一个函数，构成一个环形链表
//num:表示小孩的个数
//*Boy:返回该环形的链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}  //空结点
	curBoy := &Boy{} //空结点

	//判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	//循环的构建环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		//构成环形链表，需要一个辅助指针

		//因为第一个小孩比较特殊
		if i == 1 { //第一个小孩
			first = boy //不要动
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构成环形列表
		}
	}
	return first
}

//显示单向的环形链表[遍历]
func ShowBoy(first *Boy) {
	//处理一下环形链表为空的情况
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩")
		return
	}
	//创建一个指针，帮助遍历[说明至少有一个小孩]
	curBoy := first
	for {
		fmt.Printf("小孩的编号:%d ->", curBoy.No)
		//退出的条件
		if curBoy.Next == first { //显示到最后的小孩
			break
		}
		//curBoy移动到下一个小孩
		curBoy = curBoy.Next
	}
}

//startNum 表示从第几个小孩开始数
//countNum 表示数几下
func PlayGame(first *Boy, startNum int, countNum int) {
	//空的链表，单独处理一下
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	//startNum要小于等于小孩的总数

	//需要定义一个辅助指针，帮助我们删除一个小孩
	tail := first
	for {
		if tail.Next == first { //说明tail到了最后的小孩
			break
		}
		tail = tail.Next
	}

	//让first移动到startNo [后面删除小孩的时候，就以first为准]
	for i := 1; i <= startNum-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	//开始数countNum，然后就删除first指向的小孩
	for {
		//开始数countNum-1
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号是：%d的出圈\n", first.No)
		//删除first指向的小孩
		first = first.Next
		tail.Next = first
		//如果 tail=first 说明圈中只有一个小孩 留下这个小孩
		if tail == first {
			break
		}
	}
	fmt.Printf("最后出圈的小孩编号：%d \n", first.No)
}

//约瑟夫问题
func main() {
	first := AddBoy(5)
	//显示小孩
	ShowBoy(first)

	PlayGame(first, 2, 3)
}
