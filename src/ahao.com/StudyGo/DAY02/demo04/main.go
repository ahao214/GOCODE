package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3}
	p := &s[2] //*int，获取底层数组元素指针
	*p += 100
	fmt.Println(s)

	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data)

	//可直接修改 struct array/slice 成员
	d := [5]struct {
		x int
	}{}
	f := d[:]
	d[1].x = 10
	f[2].x = 20
	fmt.Println(d)
	fmt.Printf("%p,%p\n", &d, &d[0])

	//用append内置函数操作切片（切片追加）
	var one = []int{1, 2, 3}
	fmt.Printf("slice one:%v\n", one)
	var two = []int{4, 5, 6}
	fmt.Printf("slice two:%v\n", two)
	three := append(one, two...)
	fmt.Printf("slice three:%v\n", three)
	four := append(three, 7)
	fmt.Printf("slice four:%v\n", four)
	five := append(four, 8, 9, 10)
	fmt.Printf("slice five:%v\n", five)

	//向 slice 尾部添加数据，返回新的 slice 对象
	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s1)

	s2 := append(s1, 1)
	fmt.Printf("%p\n", s2)

	fmt.Println(s1, s2)

	//拷贝切片
	a1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice a1:%v\n", a1)

	a2 := make([]int, 10)
	fmt.Printf("slice a2:%v\n", a2)

	copy(a2, a1)
	fmt.Printf("slice a1:%v\n", a1)
	fmt.Printf("slice a2:%v\n", a2)
	a3 := []int{1, 2, 3}
	fmt.Printf("slice a3:%v\n", a3)
	a3 = append(a3, a2...)
	fmt.Printf("slice a3:%v\n", a3)
	a3 = append(a3, 4, 5, 6)
	fmt.Printf("slice a3:%v\n", a3)

	//函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。
	// 两个 slice 可指向同一底层数组，允许元素区间重叠。
	date := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array date:", date)
	e1 := date[8:]
	e2 := date[:5]
	fmt.Printf("slice e1:%v\n", e1)
	fmt.Printf("slice e2:%v\n", e2)
	copy(e2, e1)
	fmt.Printf("slice e1:%v\n", e1)
	fmt.Printf("slice e2:%v\n", e2)
	fmt.Println("last array date:", date)

	//遍历
	dataone := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := dataone[:]
	for index, value := range slice {
		fmt.Printf("index:%v,value:%v\n", index, value)
	}

	//字符串和切片
	strWorld := "hello world"
	strone := strWorld[0:5]
	fmt.Println(strone)

	strtwo := strWorld[6:]
	fmt.Println(strtwo)

	strThree := []byte(strWorld)
	strThree[6] = 'G'
	strThree = strThree[:8]
	strThree = append(strThree, '!')
	strWorld = string(strThree)
	fmt.Println(strWorld)

}
