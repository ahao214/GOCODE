package main

import (
	"fmt"
)

type shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	areaInfo := sq1
	fmt.Printf("面积为：%f\n", areaInfo.Area())
}
