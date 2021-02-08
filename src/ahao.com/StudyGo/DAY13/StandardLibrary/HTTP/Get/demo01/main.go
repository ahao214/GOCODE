package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//GET请求示例
func main() {
	resp, err := http.Get("https://www.51mh.com/")
	if err != nil {
		fmt.Println("get failed,err:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed,err:", err)
		return
	}
	fmt.Println(string(body))

}
