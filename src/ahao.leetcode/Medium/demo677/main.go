package main

//677. 键值映射
type MapSum struct {
	m, pre map[string]int
}

func Constructor() MapSum {
	return MapSum{map[string]int{}, map[string]int{}}
}

func (this *MapSum) Insert(key string, val int) {
	delta := val
	if this.m[key] > 0 {
		delta -= this.m[key]
	}
	this.m[key] = val
	for i := range key {
		this.pre[key[:i+1]] += delta
	}
}

func (this *MapSum) Sum(prefix string) int {
	return this.pre[prefix]
}
