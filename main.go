package main

import (
	day3concurency "glp/day3_concurency"
	"sync"
)

func main() {
	const requests = 10
	requestIDs := make(chan string)
	var wg sync.WaitGroup

	go day3concurency.RequestGenerator(requestIDs, requests)

	for requestID := range requestIDs {
		apiResponse := make(chan string, 1)
		wg.Add(1)
		go day3concurency.ProcessRequest(requestID, apiResponse, &wg)
	}
	wg.Wait()
}
