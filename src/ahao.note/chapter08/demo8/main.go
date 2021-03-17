package main

import (
	"fmt"
	"runtime"
	"time"
)

//Goexit
func main() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		defer fmt.Println("a")

		func() {
			defer func() {
				fmt.Println("b", recover() == nil)
			}()

			func() {
				fmt.Println("c")
				runtime.Goexit()
				fmt.Println("c done")
			}()
			fmt.Println("b done.")
		}()
		fmt.Println("a done.")
	}()
	<-exit
	fmt.Println("main exit")

	for i := 0; i < 2; i++ {
		go func(x int) {
			for n := 0; n < 2; n++ {
				fmt.Printf("%c:%d\n", 'a'+x, n)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}
	runtime.Goexit()
	fmt.Println("main exit.")
}
