package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//Hash表
func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed:", err)
		return
	}
	defer c.Close()
	_, err = c.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Ints(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}