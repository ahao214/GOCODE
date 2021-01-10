package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//将f:/one.txt里面的内容导入到f:/two.txt文件里面

	//1.首先将f:/one.txt里面的内容读取到内存
	//2.将读取到的内容写入到f:/two.txt文件里面
	filePathone := "f:/one.txt"
	filePathtwo := "f:/two.txt"
	data, err := ioutil.ReadFile(filePathone)
	if err != nil {
		fmt.Printf("read file err:%v", err)
		return
	}
	err = ioutil.WriteFile(filePathtwo, data, 0666)
	if err != nil {
		fmt.Printf("write file err=%v", err)
	}

}
