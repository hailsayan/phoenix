package main

import (
	"fmt"
)

type DNF struct {
	version string
}

func (dnf *DNF) update(pkg string) {
	fmt.Printf("updating package %s\n", pkg)
}

func main() {

	fedora := &DNF{version: "41"}
	fmt.Printf("%p\n", fedora)

	fmt.Println(fedora.version)

	fedora.update("curl")
}
