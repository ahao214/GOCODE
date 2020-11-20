package main

import "fmt"

//构造函数
type dog struct{
	name string
}


type person struct{
	name string
	age int
}

func newDog(name string) dog{
	return dog{
		name:name,
	}
}

func newPerson(name string,age int)person{
	return person{
		name:name,
		age:age,
	}
}

func(d dog) wang(){
	fmt.Printf("%s:汪汪汪",d.name)
}

//使用值接受者
//传拷贝进去
func(p person) guonian(){
	p.age++
}


//使用指针接受者
//传内存地址进去
func(p *person) zhenguonian(){
	p.age++
}

func main(){
	d1:=newDog("旺财")
	d1.wang()
	fmt.Println(d1)

	p1:=newPerson("元帅",18)
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)

	p1.zhenguonian()
	fmt.Println(p1.age)

}