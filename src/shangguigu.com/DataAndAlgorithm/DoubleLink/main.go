package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //这个表示指向前一个结点
	next     *HeroNode //这个表示指向下一个结点
}

//给链表插入一个结点
//编写第一种插入方式，在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHero *HeroNode) {
	//思路
	//1.先找到该链表的最后这个结点
	//2.创建一个辅助结点
	temp := head
	for {
		if temp.next == nil { //表示到最后
			break
		}
		temp = temp.next //让temp不断的指向下一个结点
	}
	//3.将newHero加入到链表的最后
	temp.next = newHero
	newHero.pre = temp
}

//给链表插入一个结点
//编写第一种插入方式 根据no的编号，从小到大进行插入
func InsertHeroNodeByNo(head *HeroNode, newHero *HeroNode) {
	//1.先找到适当的结点
	//2.创建一个辅助结点
	temp := head
	flag := true
	//让插入的结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil { //表示到链表的最后
			break
		} else if temp.next.no > newHero.no {
			//说明newHero就应该插入到temp的后面
			break
		} else if temp.next.no == newHero.no {
			//说明我们的链表中已经有这个no，就不能插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，已经存在no=", newHero.no)
		return
	} else {
		newHero.next = temp.next
		newHero.pre = temp
		if temp.next != nil {
			temp.next.pre = newHero
		}
		temp.next = newHero
	}

}

//删除
func DeleteHero(head *HeroNode, id int) {
	temp := head
	flag := false
	//找到要删除的结点，和temp的下一个结点的no比较
	for {
		if temp.next == nil { //表示到链表的最后
			break
		} else if temp.next.no == id {
			//说明我们找到了
			flag = true
			break
		}
		temp = temp.next
	}
	if flag { //确实找到了
		temp.next = temp.next.next //赋值之后，temp.next可能为nil
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("你要删除的id不存在")
	}
}

//显示链表的所有结点信息
func ListHeroNode(head *HeroNode) {
	//创建一个辅助结点
	temp := head
	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("链表是空的，没有数据显示")
		return
	}
	//遍历链表
	for {
		fmt.Printf("[%d %s %s]==>", temp.next.no, temp.next.name, temp.next.nickname)

		//判断是否到了链表的最后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

//显示链表的所有结点信息
func ListHeroNode2(head *HeroNode) {
	//创建一个辅助结点
	temp := head

	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("链表是空的，没有数据显示")
		return
	}

	//让temp定位到双向链表的最后结点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	//遍历链表
	for {
		fmt.Printf("[%d %s %s]==>", temp.no, temp.name, temp.nickname)

		//判断是否到了链表的头部
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

//单链表
func main() {
	//1.先创建一个头结点
	head := &HeroNode{}
	//2.创建一个新的hero结点
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
		next:     nil,
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
		next:     nil,
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
		next:     nil,
	}
	//加入
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)

	//显示
	ListHeroNode(head)
	fmt.Println("逆序打印")
	//逆序打印
	ListHeroNode2(head)
	////删除
	//DeleteHero(head, 2)
	//ListHeroNode(head)
}
