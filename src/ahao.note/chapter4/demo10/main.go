package main

import (
	"log"
)

//连续调用panic，仅最后一个会被recover捕获
func main() {
	defer func() {
		for {
			if err := recover(); err != nil {
				log.Println(err)
			} else {
				log.Fatalln("fatal")
			}
		}
	}()
	defer func() {
		panic("you are dead")
	}()
	panic("i am dead")
}
