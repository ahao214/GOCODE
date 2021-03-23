package main

import (
	"fmt"
	"reflect"
)

type Movie struct {
	Name   string
	Rating float32
}

type Address struct {
	Number int
	Street string
	City   string
}

type Superhero struct {
	Name    string
	Age     int
	Address Address
}

type Alarm struct {
	Time  string
	Sound bool
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: false,
	}
	return a
}

type Drink struct {
	Name string
	Ice  bool
}

//结构体
func main() {
	m := Movie{
		Name:   "Citizen Kan",
		Rating: 10,
	}
	fmt.Println(m.Name, m.Rating)

	e := Superhero{
		Name: "Batman",
		Age:  32,
		Address: Address{
			Number: 1007,
			Street: "Mountain Driver",
			City:   "Gotham",
		},
	}
	fmt.Printf("%+v\n", e)

	fmt.Printf("%+v\n", NewAlarm("07:00"))

	a := Drink{
		Name: "Lemon",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonada",
		Ice:  true,
	}
	if a == b {
		fmt.Println("a and b are the same")
	}
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
