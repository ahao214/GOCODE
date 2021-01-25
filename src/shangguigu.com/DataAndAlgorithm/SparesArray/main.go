package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

//稀疏数组
func main() {
	//1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //表示黑子
	chessMap[2][3] = 2 //表示蓝子

	//2.输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3.转成稀疏数组
	//思路
	//(1).遍历chessMap，如果发现有一个元素不等于，创建一个结构体
	//(2).将其放到对应的切片中

	//定义一个切片
	var sparesArr []ValNode
	//标准的一个稀疏数组，还有一个 记录元素的数组的规模(行和列，默认值)
	//创建一个ValNode值节点
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparesArr = append(sparesArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个ValNode值节点
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparesArr = append(sparesArr, valNode)
			}
		}
	}

	//输出稀疏数组
	fmt.Println("当前的稀疏数组是：")
	for i, v := range sparesArr {
		fmt.Printf("%d:%d %d %d\n", i, v.row, v.col, v.val)
	}

	//如何恢复原始的数组

	//使用稀疏数组恢复

	//先创建一个原始数组
	var chessMap2 [11][11]int
	//遍历sparesArr
	for i, valNode1 := range sparesArr {
		if i != 0 { //跳过第一行
			chessMap2[valNode1.row][valNode1.col] = valNode1.val
		}

	}
	//恢复后的原始数据
	fmt.Println("恢复后的原始数据是：")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
