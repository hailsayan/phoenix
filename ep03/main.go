package main

import "fmt"

type animal struct {
	age int
}

func (a animal) walk() {
	fmt.Println("walking")
}

type person struct {
	animal
	name string
}

func (p person) talk() {
	fmt.Println("talking")
}

func main() {
	psyon := person{
		animal: animal{
			age: 21,
		},
		name: "psyon",
	}
	_ = psyon.age
}
