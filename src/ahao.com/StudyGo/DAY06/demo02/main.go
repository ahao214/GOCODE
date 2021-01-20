package main

import (
	"fmt"
)

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, "closed")
}

func Close(t Test) {
	t.Close()
}

func main() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.Close()
	}

	for _, t := range ts {
		defer Close(t)
	}

}
