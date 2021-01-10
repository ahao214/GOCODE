package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//自己编写一个函数，它接收两个文件路径
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	defer srcFile.Close()
	//通过srcfile，获取到Reader
	reader := bufio.NewReader(srcFile)
	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	//通过dstFile，获取writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)

}

func main() {
	//将E盘下面的kebi.jpg拷贝到F盘下面

	srcFile := "e:/kebi.jpg"
	dstFile := "f:/kebi.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷贝成功")
	} else {
		fmt.Printf("错误信息是：%v\n", err)
	}
}
