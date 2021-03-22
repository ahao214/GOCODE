package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"www.golang.org",
	"www.google.com",
	"www.qq.com",
}

//WaitGroup
func main() {
	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Status)
			}
		}(url)
	}
	//等待所有请求结束
	wg.Wait()
}
