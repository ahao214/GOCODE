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

//约瑟夫问题
func main() {
	first := AddBoy(5)
	//显示小孩
	ShowBoy(first)
}
