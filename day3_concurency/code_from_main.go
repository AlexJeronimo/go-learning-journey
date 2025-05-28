package day3concurency

/* var wgWordCount sync.WaitGroup
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
resultWg.Wait() */

/* ch := make(chan int)

go day3concurency.ToChan(10, ch)

time.Sleep(1 * time.Second)

val := <-ch

fmt.Println(val) */

/* ch := make(chan int, 1)

ch <- 1

go func() {
	val := <-ch
	fmt.Println(val)
}()

time.Sleep(10 * time.Millisecond)

ch <- 2
val2 := <-ch

fmt.Println(val2) */

/* rand.Seed(time.Now().UnixNano())

ch := make(chan int, 5)
var wg sync.WaitGroup
for i := 0; i < 3; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		numToSend := rand.Intn(5) + 1
		for k := 0; k < numToSend; k++ {
			val := rand.Intn(100)
			ch <- val
		}
	}()
}

go func() {
	wg.Wait()
	close(ch)
}()
for val := range ch {
	fmt.Println(val)
} */

/* pingCh := make(chan string)

pongCh := make(chan string)

var wg sync.WaitGroup

wg.Add(2)

go Player1(pingCh, pongCh, &wg)
go Player2(pingCh, pongCh, &wg)

wg.Wait() */

/* const numJobs = 100
const numWorkers = 5

jobs := make(chan int, numWorkers)
results := make(chan int, numWorkers)
finalResultChan := make(chan float64, 1)

var wgWorkerPool sync.WaitGroup
var wgResultsCollector sync.WaitGroup

go func() {
	for i := 0; i < numJobs; i++ {
		jobs <- rand.Intn(1000)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
	}
	close(jobs)
}()

for i := 1; i <= numWorkers; i++ {
	wgWorkerPool.Add(1)
	go day3concurency.Workers(i, jobs, results, &wgWorkerPool)
}

wgResultsCollector.Add(1)
go day3concurency.ResultsCollector(results, finalResultChan, &wgResultsCollector)

wgWorkerPool.Wait()

close(results)

wgResultsCollector.Wait()

finalAverage := <-finalResultChan
fmt.Println(finalAverage) */

/* var wg sync.WaitGroup

results := make(chan string, 5)

const numWorkers = 3

for worker := range numWorkers {
	wg.Add(1)
	go day3concurency.Worker1(worker, results, &wg)
}

go func() {
	wg.Wait()
	close(results)
}()

for result := range results {
	fmt.Println(result)
}

fmt.Println("All workers completed their jobs and all results received. Programm finish...") */
