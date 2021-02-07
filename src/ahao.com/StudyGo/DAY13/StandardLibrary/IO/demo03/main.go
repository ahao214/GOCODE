package demo03

import (
	"fmt"
	"os"
)

//写文件
func main() {
	//新建文件
	file, err := os.Create("./xx.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
	}
}
