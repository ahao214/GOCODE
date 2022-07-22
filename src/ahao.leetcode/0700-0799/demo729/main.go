package main

/*
729. 我的日程安排表 I
*/

//方法一：直接遍历
type pair struct{ start, end int }
type MyCalendar []pair

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start, end int) bool {
	for _, p := range *c {
		if p.start < end && start < p.end {
			return false
		}
	}
	*c = append(*c, pair{start, end})
	return true
}


//方法二：二分查找
type MyCalendar2 struct {
	*redblacktree.Tree
}

func Constructor2() MyCalendar2 {
	t := redblacktree.NewWithIntComparator()
	t.Put(math.MaxInt32, nil) // 哨兵，简化代码
	return MyCalendar2{t}
}

func (c MyCalendar) Book2(start, end int) bool {
	node, _ := c.Ceiling(end)
	it := c.IteratorAt(node)
	if !it.Prev() || it.Value().(int) <= start {
		c.Put(start, end)
		return true
	}
	return false
}
