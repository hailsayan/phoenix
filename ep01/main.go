package main

import "fmt"

func main() {

	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")

	default:
		fmt.Println("non of the above cases")
	}

	printMyType(i)
	printMyType(true)
	printMyType("hello")

}

func printMyType(i interface{}) {
	switch i.(type) {
	case bool:
		fmt.Println("I am a bool")
	case int:
		fmt.Println("I am an int")
	default:
		fmt.Println("I am neither a bool nor an int")
	}
}
