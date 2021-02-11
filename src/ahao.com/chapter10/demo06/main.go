package main

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "纽约"
	ch <- "华盛顿"
	ch <- "伦敦"
	ch <- "北京"
	ch <- "东京"
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s", input)
	}
}

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
}
