package day3concurency

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		processedResult := job * 2

		results <- processedResult
	}
}
