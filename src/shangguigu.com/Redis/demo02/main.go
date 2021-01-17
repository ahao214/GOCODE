package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go操作redis中的Hash数据类型
	//连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭

	//通过go向hash写入数据
	_, err = conn.Do("HSet", "user", "name", "Tom")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	_, err = conn.Do("HSet", "user", "age", 20)
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	//通过redis从hash读取数据
	name, err := redis.String(conn.Do("HGet", "user", "name"))
	if err != nil {
		fmt.Println("conn.Do HGet err=", err)
		return
	}
	fmt.Println("姓名是：", name)
	age, err := redis.Int(conn.Do("HGet", "user", "age"))
	if err != nil {
		fmt.Println("conn.Do HGet err=", err)
		return
	}
	fmt.Println("年龄是：", age)
}
