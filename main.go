package main

import (
	"fmt"
	day3concurency "glp/day3_concurency"
	"sync"
)

func main() {
	filename := "app.log"
	/* count := 1000
	err := day3concurency.GenerateLogEntries(filename, count)
	if err != nil {
		fmt.Println(err)
	} */

	logLines := make(chan string, 100)
	errorLogs := make(chan string, 50)
	var wg sync.WaitGroup

	go day3concurency.ReadLogEntries(filename, logLines)

	workerPool := 5

	wg.Add(5)
	for worker := range workerPool {
		go day3concurency.LogAnalyzer(worker, logLines, errorLogs, &wg)
	}

	go func() {
		for err := range errorLogs {
			fmt.Println(err)
		}
	}()

	wg.Wait()
	close(errorLogs)

}
