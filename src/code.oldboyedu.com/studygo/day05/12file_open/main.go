package main


import(
	"fmt"
	"os"
	"io"
	"bufio"
	"ioutil"
)

func readFromFileOne(){
	fileObj,err:=os.Open("./main.go")
	if err!=nil{
		fmt.Printf("open file failed,err:%v",err)
		return 
	}
	//关闭文件
	defer fileObj.Close()

	//读取文件
	//var tmp=make([]byte,128)
	var tmp [128]byte
	for{
		n,err:=fileObj.Read(tmp[:])
		if err == io.EOF{
			fmt.Println("读完了")
			return 
		}
		if err!=nil{
			fmt.Printf("read from file failed,err:%v",err)
			return
		}
		fmt.Printf("读取了%d个字节\n",n)
		fmt.Println(string(tmp[:n]))
		if n<128{
			return
		}
	}
}

func readFromFilebyBufio(){
	fileObj,err:=os.Open("./main.go")	
	if err!=nil{
		fmt.Printf("open file failed,err:%v",err)
		return
	}
	defer fileObj.Close()

	reader:=bufio.NewReader(fileObj)
	for{
		line,err:=reader.ReadString('\n')
	if err == io.EOF{
		return 
	}
	if err!=nil{
		fmt.Printf("read line failed,err:%v",err)
		return
	}
	fmt.Println(line)
	}
}

func readFromFileByIoutil(){
	ret,err:=ioutil.ReadFile("./main.go")
	if err!=nil{
		fmt.Printf("read file failed,err:%v",err)
		return
	}
	fmt.Println(string(ret))


}

func main(){
	//readFromFilebyBufio()
	readFromFileByIoutil()


}