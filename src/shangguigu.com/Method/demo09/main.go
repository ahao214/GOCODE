package main

import(
	"fmt"
)

type Box struct{
	len float64
	width float64
	height float64
}

func(box *Box)getVol()float64{
	return box.len*box.width*box.height
}

func main(){
	var box Box
	box.len=1.1
	box.width=1.1
	box.height=2.0

	vol:=box.getVol()
	fmt.Printf("体积为：%.2f",vol)
	

}