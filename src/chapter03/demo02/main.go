package main

import(
	"fmt"
)
func main(){
	m:=map[int]int{1:1,2:4,3:9}
	for k,v:=range(m){
		fmt.Printf("key:%d value:%d\n",k,v)
	}

}