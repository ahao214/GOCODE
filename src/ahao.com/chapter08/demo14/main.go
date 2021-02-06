package main

import (
	"fmt"
)

type HttpRespost struct {
	status_code int
}

func (r *HttpRespost) validResponse() {
	r.status_code = 200
}

func (r HttpRespost) updateResponse() string {
	return fmt.Sprint(r)
}

func main() {
	var rone HttpRespost //rone是值
	rone.validResponse()
	fmt.Println(rone.updateResponse())

	rtwo := new(HttpRespost) //rtwo是指针
	rtwo.validResponse()
	fmt.Println(rtwo.updateResponse())

}
