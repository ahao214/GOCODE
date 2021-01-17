package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//链接池

//定义一个全局的pool
var pool *redis.Pool

//当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接的代码，链接哪个ip
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//先从pool，取出一个链接
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", "name", "汤姆猫")
	if err != nil {
		fmt.Println("conn.Do set err=", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do get err=", err)
		return
	}
	fmt.Println("r=", r)

	//如果我们从pool中取出链接，一定要保证链接池没有关闭

}
