package main

import (
	"fmt"
	day3concurency "glp/day3_concurency"
	"sync"
)

func main() {
	const numProducts = 10

	productRequests := make(chan day3concurency.ProductRequest, 100)
	partialDataChan := make(chan day3concurency.ProductPartialData, 100)
	finalProductsChan := make(chan day3concurency.Product, numProducts)

	var wgGenerators sync.WaitGroup // Для генератора запитів
	var wgFetchers sync.WaitGroup   // Для всіх воркерів FetchPrice, FetchDescription, FetchAvailability
	var wgAggregator sync.WaitGroup // Для агрегатора

	wgGenerators.Add(1)
	go day3concurency.GenerateProductRequests(productRequests, numProducts)
	wgFetchers.Add(3)

	priceRequests := make(chan day3concurency.ProductRequest, 100)
	descRequests := make(chan day3concurency.ProductRequest, 100)
	availRequests := make(chan day3concurency.ProductRequest, 100)

	go day3concurency.FetchPrice(priceRequests, partialDataChan, &wgFetchers)
	go day3concurency.FetchDescription(descRequests, partialDataChan, &wgFetchers)
	go day3concurency.FetchAvailability(availRequests, partialDataChan, &wgFetchers)

	wgAggregator.Add(1)

	pricePartialDataChan := make(chan day3concurency.ProductPartialData, numProducts)
	descPartialDataChan := make(chan day3concurency.ProductPartialData, numProducts)
	availPartialDataChan := make(chan day3concurency.ProductPartialData, numProducts)

	go day3concurency.AggregateProductData(pricePartialDataChan, descPartialDataChan, availPartialDataChan, finalProductsChan, numProducts, &wgAggregator)
	wgGenerators.Add(1)
	go day3concurency.GenerateProductRequests(productRequests, numProducts)

	wgFetchers.Add(3 * numProducts)
	go func() {
		defer close(pricePartialDataChan)
		defer close(descPartialDataChan)
		defer close(availPartialDataChan)

		// Чекаємо, доки генератор запитів завершить роботу.
		wgGenerators.Wait()

		makeReqChan := func(req day3concurency.ProductRequest) <-chan day3concurency.ProductRequest {
			ch := make(chan day3concurency.ProductRequest, 1)
			ch <- req
			close(ch)
			return ch
		}

		for req := range productRequests { // Отримуємо запити від генератора
			// Запускаємо окремі Goroutine для кожного типу Fetch-операції для поточного запиту.
			// Кожна з них додає 1 до wgFetchers.
			wgFetchers.Add(1)
			go day3concurency.FetchPrice(makeReqChan(req), pricePartialDataChan, &wgFetchers) // makeReqChan створює канал з одним запитом
			wgFetchers.Add(1)
			go day3concurency.FetchDescription(makeReqChan(req), descPartialDataChan, &wgFetchers)
			wgFetchers.Add(1)
			go day3concurency.FetchAvailability(makeReqChan(req), availPartialDataChan, &wgFetchers)
		}
		fmt.Println("Main: Всі productRequests розподілені серед воркерів.")
	}()

	wgAggregator.Add(1)
	go day3concurency.AggregateProductData(pricePartialDataChan, descPartialDataChan, availPartialDataChan, finalProductsChan, numProducts, &wgAggregator)

	wgFetchers.Wait()
	wgAggregator.Wait()

	countFinalProducts := 0
	for p := range finalProductsChan { // Читаємо з finalProductsChan, поки він не закриється
		fmt.Printf("Продукт ID: %s, Ціна: %.2f, Опис: '%s', Наявність: %t\n",
			p.ID, p.Price, p.Description, p.Availability)
		countFinalProducts++
	}
	fmt.Printf("Main: Зібрано всього %d агрегованих продуктів.\n", countFinalProducts)
	fmt.Println("Main: Програма завершена.")
}
