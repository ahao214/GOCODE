package main

//Queue和Set都直接实现，Set实现了线程安全，Queue不是安全的线程
import (
	"fmt"
	"github.com/isdamir/gotype"
)

type LRU struct {
	cacheSize int
	queue     *gotype.SliceQueue
	hashSet   *gotype.Set
}

//判断缓存队列是否已满
func (p *LRU) IsQueueFull() bool {
	return p.queue.Size() == p.cacheSize
}

//把页号为pageNum的页页缓存到队列中，同时也添加到Hash表中
func (p *LRU) EnQueue(pageNum int) {
	//如果队列满了，需要删除队尾的缓存的页
	if p.IsQueueFull() {
		p.hashSet.Remove(p.queue.PopBack())
	}
	p.queue.EnQueueFirst(pageNum)
	//把新缓存的结点同时添加到hash表中
	p.hashSet.Add(pageNum)
}

//当访问某一个page的时候会调用这个函数，对于访问的page有两种情况
//1.如果page在缓存队列中，直接把这个结点移动到队首
//2.如果page不在缓存队列中，把这个page缓存到队首
func (p *LRU) AccessPage(pageNum int) {
	//page不在缓存队列中，把它缓存到队首
	if !p.hashSet.Contains(pageNum) {
		p.EnQueue(pageNum)
	} else if pageNum != p.queue.GetFront() { //page已经在缓存队列中了，移动到队首
		p.queue.Remove(pageNum)
		p.queue.EnQueueFirst(pageNum)
	}
}

func (p *LRU) PrintQueue() {
	for !p.queue.IsEmpty() {
		fmt.Print(p.queue.DeQueue(), "")
	}
	fmt.Println()
}

//实现URL缓存方案
func main() {
	//假设缓存大小为3
	lru := &LRU{cacheSize: 3, queue: gotype.NewSliceQueue(), hashSet: gotype.NewSet()}
	lru.AccessPage(1)
	lru.AccessPage(2)
	lru.AccessPage(3)
	lru.AccessPage(4)
	lru.AccessPage(1)
	lru.AccessPage(6)
	lru.AccessPage(7)
	//通过上面的访问序列后，缓存的信息为
	lru.PrintQueue()

}
