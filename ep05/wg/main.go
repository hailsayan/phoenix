package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup

	file := "websites.txt"
	reader, _ := os.Open(file)
	scanner := bufio.NewScanner(reader) //line by line
	for scanner.Scan() {
		go getWebsite(scanner.Text(), &wg)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("total time: ", time.Since(startTime))
}

func getWebsite(website string, wg *sync.WaitGroup) {
	if res, err := http.Get(website); err != nil {
		fmt.Println(website, "is down")
	} else {
		fmt.Printf("[%d] %s is up\n", res.StatusCode, website)
	}
	wg.Done()
}
