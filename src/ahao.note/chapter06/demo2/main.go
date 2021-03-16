package main

import (
	"fmt"
	"reflect"
)

type S struct {
}
type T struct {
	S
}

func (S) sVal() {
}

func (*S) sPtr() {

}

func (T) tVal() {

}

func (*T) tPtr() {

}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)

	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

//方法集
func main() {
	var t T
	methodSet(t)
	fmt.Println("------")
	methodSet(&t)
}
