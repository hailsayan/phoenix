package main

import (
	"fmt"
	"math"
)

// type DNF struct {
// 	version string
// }

// func (dnf *DNF) update(pkg string) {
// 	fmt.Printf("updating package %s\n", pkg)
// }

type geometry interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func printArea(g geometry) {
	fmt.Printf("the geometry area is %f\n", g.area())
}

func main() {

	r1 := &rectangle{width: 5, height: 3}
	c1 := &circle{radius: 4}

	printArea(r1)
	printArea(c1)

	// fedora := DNF{version: "41"}
	// fmt.Printf("%p\n, &fedora")

	// fmt.Println(fedora.version)

	// fedora.update("curl")
}
