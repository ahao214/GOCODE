package main

import(
	"fmt"
	"os"
	"bufio"
	"io"
)

func main(){

	//打开文件
	file,err:=os.Open("f:/one.txt")
	if err!=nil{
		fmt.Println("open file err=",err)
	}
	fmt.Printf("file=%v",file)
	

	//当函数退出时，要及时关闭file
	defer file.Close()		//要及时关闭file，否则会内存泄漏

	//创建一个*Reader，是带缓存的
	reader:=bufio.NewReader(file)
	//循环读取文件的内容
	for{
		str,err:=reader.ReadString('\n')	//读到一个换行就结束
		if err==io.EOF{
			//io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")


}