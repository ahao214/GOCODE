package main

import (
	"fmt"
)

//定义猫的结构体
type CatNode struct {
	no   int    //猫的编号
	name string //猫的名字
	next *CatNode
}

//添加
func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //构成一个环状
		fmt.Println(newCatNode, "加入到环形链表中")
		return
	}
	//定义一个临时的变量
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

//输出环形的链表
func ShowList(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空的链表")
		return
	}
	for {
		fmt.Printf("猫的信息是[id=%d,name=%s]：\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//删除
func DeleteCat(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表")
		return head
	}
	//如果只有一个结点
	if temp.next == head {
		temp.next = nil
		return head
	}
	//将helper 定位到链表的最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//如果有两个及以上的结点
	flag := true
	for {
		if temp.next == head { //如果到这里，说明我比较到最后一个【最后一个还没比较】
			break
		}
		if temp.no == id {
			//恭喜找到 可以在这里删除
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}
		temp = temp.next     //移动【比较】
		helper = helper.next //移动 【一旦找到要删除的结点 helper】
	}
	//这里还要比较一次
	if flag { //如果flag为真，则我们上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Printf("对不起，没有No=%d的猫猫\n", id)
		}
	}
	return head
}

//单向环形链表
func main() {
	//这里我们初始化一个环形链表的头结点
	head := &CatNode{}
	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "Tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "Jack",
	}
	cat3 := &CatNode{
		no:   3,
		name: "Kevin",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)

	fmt.Println("环形链表的情况如下：")
	ShowList(head)

	head = DeleteCat(head, 3)

}
