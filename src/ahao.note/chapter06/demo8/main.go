package main

import (
	"fmt"
)

func main() {
	var cheeses [2]string
	cheeses[0] = "Marilll"
	cheeses[1] = "Epos"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)

	var players = make(map[string]int)
	players["cook"] = 32
	players["LBL"] = 36
	players["stokes"] = 22
	fmt.Println(players["cook"])
	delete(players, "cook")

}
