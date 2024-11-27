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

	// jobs
	jobs := make(chan string, 5)

	// worker
	const workers = 3
	for worker := 0; worker < workers; worker++ {
		go getWebsite(jobs, &wg, worker)
	}

	//job producer
	file := "websites.txt"
	reader, _ := os.Open(file)
	scanner := bufio.NewScanner(reader) //line by line
	for scanner.Scan() {
		jobs <- scanner.Text()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("total time: ", time.Since(startTime))
}

func getWebsite(websites chan string, wg *sync.WaitGroup, worker int) {
	for website := range websites {
		if res, err := http.Get(website); err != nil {
			fmt.Printf("%s is down, processed by worker %d", website, worker)
		} else {
			fmt.Printf("[%d] %s is up, processed by worker %d\n", res.StatusCode, website, worker)
		}
		wg.Done()
	}

}
