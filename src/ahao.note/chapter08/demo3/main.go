package main

import (
	"fmt"
	"sync"
	"time"
)

//sync.WaitGroup
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println("goroutine,", id, "done.")
		}(i)
	}
	fmt.Println("main...")
	wg.Wait()
	fmt.Println("main exit.")
}
