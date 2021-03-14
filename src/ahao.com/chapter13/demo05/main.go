package main

import (
	"log"
	"os"
)

//文件处理-删除
func main() {
	err := os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}
