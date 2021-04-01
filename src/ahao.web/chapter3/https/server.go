package main

import "net/http"

//带有附加配置的Web服务器
func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
