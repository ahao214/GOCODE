package main

import(
	"fmt"
)

func main(){
	var a map[string]string
	a = make(map[string]string,10)
	a["n1"]="宋江"
	a["n2"]="李逵"
	a["n1"]="武松"
	a["n3"]="林冲"
	fmt.Println(a)


	cities:=make(map[string]string)
	cities["1"]="上海"
	cities["2"]="北京"
	cities["3"]="广州"
	cities["4"]="深圳"
	fmt.Println(cities)

	heros:=map[string]string{
		"h1":"钢铁侠",
		"h2":"蝙蝠侠",
		"h3":"蜘蛛侠",
	}
	heros["h4"]="豆豆侠"
	fmt.Println("heros:",heros)

}