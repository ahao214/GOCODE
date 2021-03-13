package main

import (
	"log"
	"os"
)

//文件处理-重命名与移动
func main() {
	originalPath := "test.txt"
	newPath := "test2.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
