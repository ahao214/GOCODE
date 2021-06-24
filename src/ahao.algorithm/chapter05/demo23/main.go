package main

import (
	"container/list"
	"fmt"
	. "github.com/isdamir/gotype"
)

type QItem struct {
	word string
	len  int
}

func NewQItem(word string, len int) *QItem {
	return &QItem{word, len}
}

//判断两个字符串是否只要一个不同的字符
func isadjacent(a, b string) bool {
	diff := 0
	for i, v := range a {
		if byte(v) != b[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	if diff == 1 {
		return true
	} else {
		return false
	}
}

//返回从start到target的最短链
func shortestChainLen(start, target string, D *list.List) int {
	Q := NewSliceQueue()
	item := NewQItem(start, 1)
	Q.EnQueue(item)
	for !Q.IsEmpty() {
		cur := Q.DeQueue().(*QItem)
		var n *list.Element
		for e := D.Front(); e != nil; e = n {
			n = e.Next()
			//如果这两个字符串只有一个字符不同
			tmp := e.Value.(string)
			if isadjacent(cur.word, tmp) {
				item.word = tmp
				item.len = cur.len + 1
				Q.EnQueue(item) //把这个字符串放入到队列中
				//把这个字符串从队列中删除来避免被重复遍历
				D.Remove(e)
				//通过转变后得到了目标字符
				if tmp == target {
					return item.len
				}
			}
		}
	}
	return 0
}

//查找到达目标词的最短链长度
func main() {
	D := list.New()
	D.PushBack("pooN")
	D.PushBack("pbcc")
	D.PushBack("zamc")
	D.PushBack("poIc")
	D.PushBack("pbca")
	D.PushBack("pbIc")
	D.PushBack("poIN")
	start := "TooN"
	target := "pbca"
	fmt.Println("最短的链条的长度为：", shortestChainLen(start, target, D))
}
