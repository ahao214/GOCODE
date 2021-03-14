package main

import (
	"fmt"
	"log"
	"os"
)

var (
	file     *os.File
	err      error
	dirPath  = "test/test2/"
	fileName = "demo"
	filePath = dirPath + fileName
)

//文件处理-跳转函数
func main() {
	file, _ := os.Open(filePath)
	defer file.Close()
	//偏离位置，可以是正数也可以是负数
	var offset int64 = 5
	//用来计算offset的初始位置
	//0-文件开始位置
	//1-当前位置
	//2-文件结尾处
	whence := 0
	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("移动到位置5：", newPosition)
	//从当前位置回退两字节
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("从当前位置退回两字节：", newPosition)
	//使用下面的方式得到当前的位置
	currentPosition, err := file.Seek(0, 1)
	fmt.Println("当前位置：", currentPosition)
	//转到文件开始处
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("转到文件开始位置：(0,0)", newPosition)

}
