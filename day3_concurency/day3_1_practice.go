package day3concurency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ToChan(n int, ch chan int) {
	ch <- n
	time.Sleep(1 * time.Second)
}

//---PING PONG---

func Player1(pingCh chan string, pongCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		pingCh <- "Ping"

		msg := <-pongCh
		fmt.Println(msg)

		time.Sleep(100 * time.Millisecond)
	}

	close(pingCh)

}

func Player2(pingCh chan string, pongCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		msg, ok := <-pingCh
		if !ok {
			break
		}

		fmt.Println(msg)

		pongCh <- "Pong"
	}
}

//----------------

func Workers(in int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		results <- job
	}
}

func ResultsCollector(results <-chan int, finalResultChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()

	var allResults []int
	for res := range results {
		allResults = append(allResults, res)
	}

	if len(allResults) == 0 {
		finalResultChan <- 0.0
		return
	}

	sum := 0

	for _, num := range allResults {
		sum += num
	}

	average := float64(sum) / float64(len(allResults))

	finalResultChan <- average
}
