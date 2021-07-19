package main

import (
	"bytes"
	"fmt"
	. "github.com/isdamir/gotype"
	"strconv"
)

type NumGroup struct {
	numbers     []int
	n           int
	visited     []bool
	graph       [][]int
	combination []int
	s           *Set
}

func NewnumGroup(arr []int) *NumGroup {
	return &NumGroup{
		numbers:     arr,
		n:           len(arr),
		visited:     make([]bool, len(arr)),
		graph:       make([][]int, len(arr)),
		combination: make([]int, 0),
		s:           NewSet(),
	}
}

//转化int数组为string
func iTos(a []int) string {
	var buff bytes.Buffer
	for _, v := range a {
		buff.WriteString(strconv.Itoa(v))
	}
	return buff.String()
}

//对图从结点start位置开始进行深度遍历
func (p *NumGroup) depthFirstSearch(start int) {
	p.visited[start] = true
	p.combination = append(p.combination, p.numbers[start])
	if len(p.combination) == p.n {
		//4不出现在第三个位置
		if p.combination[2] != 4 {
			p.s.Add(iTos(p.combination))
		}
	}
	for j := 0; j < p.n; j++ {
		if p.graph[start][j] == 1 && !p.visited[j] {
			p.depthFirstSearch(j)
		}
	}

	p.combination = p.combination[:len(p.combination)-1]
	p.visited[start] = false
}

//获取1、2、3、4、5的左右组合，使得"4"不能在第三位，"3"与"5"不能相连
func (p *NumGroup) getAllCombinations() {
	for i := range p.graph {
		p.graph[i] = make([]int, p.n)
	}
	//构造图
	for i := 0; i < p.n; i++ {
		for j := 0; j < p.n; j++ {
			if i == j {
				p.graph[i][j] = 0
			} else {
				p.graph[i][j] = 1
			}
		}
	}
	//确保在遍历的时候3与5是不可达的
	p.graph[3][5] = 0
	p.graph[5][3] = 0
	//分别从不同的结点出发深度遍历图
	for i := 0; i < p.n; i++ {
		p.depthFirstSearch(i)
	}
}

func (p *NumGroup) printAllCombinations() {
	res := p.s.List()
	for _, v := range res {
		fmt.Println(v)
	}
}

//求数字的组合
func main() {
	numGroup := NewnumGroup([]int{1, 2, 2, 3, 4, 5})
	numGroup.getAllCombinations()
	numGroup.printAllCombinations()
}
