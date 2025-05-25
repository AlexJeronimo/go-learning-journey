package main

import (
	"fmt"
	day3concurency "glp/day3_concurency"
	"math/rand"
	"sync"
)

func main() {
	var wgWordCount sync.WaitGroup
	text := `
	Goroutines — це легкі, незалежні потоки виконання, якими керує Go-рантайм. 
	Вони набагато дешевші та легші за звичайні системні потоки (threads). 
	Ви можете без проблем запускати сотні тисяч або мільйони goroutines у своїй програмі.

	Канали — це основний спосіб комунікації між goroutines. 
	Вони надають безпечний, синхронізований механізм для надсилання та отримання даних. 
	Ідея полягає в тому, що замість того, щоб goroutines обмінювалися даними через спільну пам'ять 
	(що може призвести до проблем з гонками даних), вони обмінюються даними шляхом передачі їх через канали.
	
	`
	wordCountChan := make(chan map[string]int)
	wgWordCount.Add(1)
	go day3concurency.CountWords(text, wordCountChan, &wgWordCount)

	value := <-wordCountChan
	for k, v := range value {
		fmt.Printf("%s: %d\n", k, v)
	}
	go func() {
		wgWordCount.Wait()
		close(wordCountChan)
	}()
	//--------------------------------------

	var wgWorkerPoll sync.WaitGroup
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int)
	results := make(chan int)

	for i := 1; i <= numWorkers; i++ {
		wgWorkerPoll.Add(1)

		go func(id int) {
			defer wgWorkerPoll.Done()
			day3concurency.Worker(id, jobs, results)
		}(i)
	}

	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- rand.Intn(100)
			fmt.Printf("Generator: send task %d\n", j)
		}
		close(jobs)
		fmt.Println("Generator: tasks channel closed")
	}()

	var resultWg sync.WaitGroup
	resultWg.Add(1)
	go func() {
		defer resultWg.Done()
		fmt.Println("Result collector: waiting for results....")
		for range numJobs {
			result := <-results
			fmt.Printf("Result collector: received result %d\n", result)
		}
		fmt.Println("Result collector: received all results.")
	}()

	wgWorkerPoll.Wait()
	resultWg.Wait()
}
