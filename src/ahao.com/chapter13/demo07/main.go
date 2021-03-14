package main

import (
	"io"
	"log"
	"os"
)

var (
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	path     = "test/text2/"
	fileName = "demo"
	filePath = path + fileName
)

//文件处理-复制文件
func main() {
	//打开原始文件
	originalFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()
	//创建新的文件作为目标文件
	newFile, err := os.Create(filePath + "_copy")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	//从源文件中复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("文件已复制，大小%dbytes", bytesWritten)
	//将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

}
