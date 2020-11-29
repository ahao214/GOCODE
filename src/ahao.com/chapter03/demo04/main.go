package main

import(
	"fmt"
	"strings"
)
var valueone float64


func main(){
	s:="abcd你"
	fmt.Println(s[4:]+"好")
	str:="你好，"+"GO 语言"
	fmt.Println(str)

	c:="我是中国人"
	for _,v:=range c{
		fmt.Printf("%c\n",v)
	}

	//替换
	strone:="你好世界，这个世界真好"
	new:="地球"
	old:="世界"
	n:=1
	m:=2
	fmt.Println(strings.Replace(strone,old,new,n))
	fmt.Println(strings.Replace(strone,old,new,m))


	//统计
	strCount:="Golang is cool,right?"
	var many="o"
	fmt.Printf("%d\n",strings.Count(strCount,many))
	fmt.Printf("%d\n",strings.Count(strCount,"oo"))
}