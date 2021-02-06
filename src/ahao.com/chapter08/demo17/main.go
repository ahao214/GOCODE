package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Abs(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	Name string
}

func main() {
	n := &NamedPoint{Point{3, 4}, "python go"}
	fmt.Println(n.Abs())
}
