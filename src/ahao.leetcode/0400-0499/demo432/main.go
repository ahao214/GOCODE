package main

import "container/list"

/*
432. 全 O(1) 的数据结构
请你设计一个用于存储字符串计数的数据结构，并能够返回计数最小和最大的字符串。

实现 AllOne 类：

AllOne() 初始化数据结构的对象。
inc(String key) 字符串 key 的计数增加 1 。如果数据结构中尚不存在 key ，那么插入计数为 1 的 key 。
dec(String key) 字符串 key 的计数减少 1 。如果 key 的计数在减少后为 0 ，那么需要将这个 key 从数据结构中删除。测试用例保证：在减少计数前，key 存在于数据结构中。
getMaxKey() 返回任意一个计数最大的字符串。如果没有元素存在，返回一个空字符串 "" 。
getMinKey() 返回任意一个计数最小的字符串。如果没有元素存在，返回一个空字符串 "" 。
*/

//双向链表+哈希表
type node struct {
	keys  map[string]struct{}
	count int
}

type AllOne struct {
	*list.List
	nodes map[string]*list.Element
}

func Constructor() AllOne {
	return AllOne{list.New(), map[string]*list.Element{}}
}

func (l *AllOne) Inc(key string) {
	if cur := l.nodes[key]; cur != nil {
		curNode := cur.Value.(node)
		if nxt := cur.Next(); nxt == nil || nxt.Value.(node).count > curNode.count+1 {
			l.nodes[key] = l.InsertAfter(node{map[string]struct{}{key: {}}, curNode.count + 1}, cur)
		} else {
			nxt.Value.(node).keys[key] = struct{}{}
			l.nodes[key] = nxt
		}
		delete(curNode.keys, key)
		if len(curNode.keys) == 0 {
			l.Remove(cur)
		}
	} else { // key 不在链表中
		if l.Front() == nil || l.Front().Value.(node).count > 1 {
			l.nodes[key] = l.PushFront(node{map[string]struct{}{key: {}}, 1})
		} else {
			l.Front().Value.(node).keys[key] = struct{}{}
			l.nodes[key] = l.Front()
		}
	}
}

func (l *AllOne) Dec(key string) {
	cur := l.nodes[key]
	curNode := cur.Value.(node)
	if curNode.count > 1 {
		if pre := cur.Prev(); pre == nil || pre.Value.(node).count < curNode.count-1 {
			l.nodes[key] = l.InsertBefore(node{map[string]struct{}{key: {}}, curNode.count - 1}, cur)
		} else {
			pre.Value.(node).keys[key] = struct{}{}
			l.nodes[key] = pre
		}
	} else { // key 仅出现一次，将其移出 nodes
		delete(l.nodes, key)
	}
	delete(curNode.keys, key)
	if len(curNode.keys) == 0 {
		l.Remove(cur)
	}
}

func (l *AllOne) GetMaxKey() string {
	if b := l.Back(); b != nil {
		for key := range b.Value.(node).keys {
			return key
		}
	}
	return ""
}

func (l *AllOne) GetMinKey() string {
	if f := l.Front(); f != nil {
		for key := range f.Value.(node).keys {
			return key
		}
	}
	return ""
}
