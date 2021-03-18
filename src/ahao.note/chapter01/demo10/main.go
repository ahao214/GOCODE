package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

//map
func main() {
	//使用字面量创建
	ma := map[string]int{"a": 1, "b": 2}
	fmt.Println(ma["a"])
	fmt.Println(ma["b"])
	//使用make创建
	one := make(map[int]string)
	two := make(map[int]string)
	one[1] = "tom"
	two[1] = "jack"
	fmt.Println(one[1])
	fmt.Println(two[1])

	thr := make(map[int]string)
	thr[1] = "tom"
	thr[1] = "tomone"
	thr[2] = "jack"
	thr[3] = "jerry"
	fmt.Println(thr[1])
	fmt.Println(len(thr))
	for k, v := range thr {
		fmt.Println("key=", k, "value=", v)

	}

	mb := make(map[int]User)
	andes := User{
		name: "不二",
		age:  18,
	}
	mb[1] = andes
	andes.age = 19
	fmt.Printf("%v\n", mb)
}
