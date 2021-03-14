package main

import (
	"log"
	"os"
)

var (
	file     *os.File
	fileInfo os.FileInfo
	err      error
	dirPath  = "test/test2/"
	fileName = "demo"
	filtPath = dirPath + fileName
)

//文件处理-打开与关闭
func main() {
	//简单的以只读方式打开
	file, err = os.Open(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	//打印文件内容
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	file.Close()
	file, err = os.OpenFile(dirPath, os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

}
