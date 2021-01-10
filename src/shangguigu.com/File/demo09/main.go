package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int //记录英文个数
	NumCount   int //记录数字个数
	SpaceCount int //记录空格个数
	OtherCount int //记录其他字符个数
}

//统计文件里面的字符个数(英文，数字，空格和其他字符)
func main() {

	fileName := "f:/char.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()

	//定义一个CharCount实例
	var count CharCount
	//创建一个reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //读取到文件的末尾
			break
		}
		//遍历str，进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透处理
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	//输出统计的结果
	fmt.Printf("字符的个数是：%v\n数字的个数是：%v\n空格的个数是：%v\n其他字符的个数是：%v\n",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}
