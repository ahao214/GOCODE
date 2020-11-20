package main

import "fmt"

//定义一个car接口类型
type car interface{
	run()
}

type falali struct{
	brand string
}

func(f falali)run(){
	fmt.Printf("%s的速度是190\n",f.brand)
}

type baoshijie struct{
	brand string
}

func(b baoshijie)run(){
	fmt.Printf("%s的速度是200\n",b.brand)
}

func drive(c car){
	c.run()
}

func main(){
	var f=falali{
		brand:"法拉利",
	}

	var b = baoshijie{
		brand:"保时捷",
	}

	drive(f)
	drive(b)
}