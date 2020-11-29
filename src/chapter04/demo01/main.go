package main

import(
	"fmt"
	"strings"
)

func main(){
	s:="A,B,C"
	parts:=strings.Split(s,",")
	for _, part:=range parts{
		fmt.Println(part)
		
	}

}