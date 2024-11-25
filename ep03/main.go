package main

import (
	"fmt"
	"time"
)

func process(num int, done chan bool) {
	for i := 0; i < num; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("process %d processed\n", i)
	}

	done <- true
}

func main() {
	done := make(chan bool, 1)

	fmt.Println("before")
	go process(4, done)
	fmt.Println("after")

	<-done
}
