package main

import (
	"fmt"
)

//编写一个函数，完成老鼠找路
//myMap 代表地图
//i和j分别表示对地图的哪个点进行测试
func FindMap(myMap *[8][7]int, i int, j int) bool {
	//分析出什么情况下，就找到出路
	if myMap[6][5] == 2 { //找到出路
		return true
	} else { //继续找
		if myMap[i][j] == 0 { //如果该点是可以探测
			//假设这个点是可以通的，但是需要探测 下右上左
			myMap[i][j] = 2

			if FindMap(myMap, i+1, j) { //向下
				return true
			} else if FindMap(myMap, i, j+1) { //向右
				return true
			} else if FindMap(myMap, i-1, j) { //向上
				return true
			} else if FindMap(myMap, i, j-1) { //向左
				return true
			} else { //死路
				myMap[i][j] = 3
				return false
			}
		} else { //说明这个点不能探测 为1，表示是墙壁
			return false
		}
	}

}

func main() {
	//先创建一个二维数组，模拟一个迷宫
	//规则
	//1.如果元素值是是1，就表示是墙壁
	//2.如果元素的值是0，是没有走过的
	//3.如果元素的值是2，表示这是一条通路
	//4.如果元素的值是3，表示是走过的一条路，但是走不通

	var myMap [8][7]int

	//先把地图的最上面和最下面设置为墙壁，数值为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}

	//把地图的最左边和最右边设置为墙壁，数值为1
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	myMap[3][1] = 1
	myMap[3][2] = 1

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	FindMap(&myMap, 1, 1)
	fmt.Println("检查完毕")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
