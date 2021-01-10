package main

import (
	"bufio"
	"fmt"
	"os"
)

//写文件操作
func main() {
	filePath := "f:/one.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=", err)
		return
	}
	defer file.Close()

	str := "hello,go语言\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为Writer是带缓存的，因此在调用WriteString方法时，其实
	//内容是先写入缓存，所以需要调用Flush()这个方法，将缓存的数据真正写入到文件中
	//否则文件中会没有数据
	writer.Flush()
}
