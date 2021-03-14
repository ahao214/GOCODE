package main

import (
	"log"
	"os"
)

//文件处理-截断
func main() {
	err := os.Truncate("test.txt", 100) //裁剪一个文件到100字节
	if err != nil {
		log.Fatal(err)
	}
}
