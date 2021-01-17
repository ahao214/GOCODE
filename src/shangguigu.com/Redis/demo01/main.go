package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis" //引入redis包
)

func main() {
	//通过go向redis写入数据和读取数据
	//1.连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()
	//2.通过go向redis写入数据string [key-value]
	_, err = conn.Do("Set", "name", "桃木")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//3.通过go向redis获取数据string [key-value]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}

	//因为返回的r是interface{}
	//因为name对应的值是string，因此需要转换

	fmt.Println("获取到的值是：", r)

	fmt.Println("操作成功")

}
