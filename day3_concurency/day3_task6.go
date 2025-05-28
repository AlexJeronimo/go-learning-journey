package day3concurency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ProductRequest struct {
	ID string
}

type ProductPartialData struct {
	ID           string
	Source       string
	Data         string
	Price        float64
	Description  string
	Availability bool
}

type Product struct {
	ID           string
	Price        float64
	Description  string
	Availability bool
}

func GenerateProductRequests(out chan<- ProductRequest, count int) {
	for n := range count {
		request := ProductRequest{
			ID: fmt.Sprintf("prod_%03d", n),
		}
		out <- request
	}
	close(out)
}

func FetchPrice(in <-chan ProductRequest, out chan<- ProductPartialData, wg *sync.WaitGroup) {
	defer wg.Done()

	for r := range in {
		timeout := rand.Intn(100) + 50
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		price := 10.0 + rand.Float64()*1000.0

		product := ProductPartialData{
			ID:     r.ID,
			Source: "PriceAPI",
			Price:  price,
		}

		out <- product
	}

}

func FetchDescription(in <-chan ProductRequest, out chan<- ProductPartialData, wg *sync.WaitGroup) {
	defer wg.Done()

	for r := range in {
		timeout := rand.Intn(100) + 50
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		descriptions := []string{
			"Високоякісний товар.",
			"Сучасний дизайн та висока функціональність.",
			"Ідеально підходить для повсякденного використання.",
			"Екологічно чисті матеріали.",
			"Обмежена серія, ексклюзивний продукт.",
		}

		product := ProductPartialData{
			ID:          r.ID,
			Source:      "DescService",
			Description: descriptions[rand.Intn(len(descriptions))],
		}

		out <- product
	}
}

func FetchAvailability(in <-chan ProductRequest, out chan<- ProductPartialData, wg *sync.WaitGroup) {
	defer wg.Done()
	for r := range in {
		timeout := rand.Intn(100) + 50
		time.Sleep(time.Duration(timeout) * time.Millisecond)

		availability := rand.Intn(2) == 1

		product := ProductPartialData{
			ID:           r.ID,
			Source:       "StockDB",
			Availability: availability,
		}

		out <- product
	}

}

func AggregateProductData(priceChan, descChan, availChan <-chan ProductPartialData, finalProducts chan<- Product, numProducts int, wg *sync.WaitGroup) {
	defer wg.Done()

	bufer := make(map[string]*Product)
	receivedPartsCount := make(map[string]int)

	closedChannels := 0
	const channelsTotal = 3

	for {
		select {
		case partialData, ok := <-priceChan:
			if !ok {
				closedChannels++
				priceChan = nil
				continue
			}
			product, exists := bufer[partialData.ID]
			if !exists {
				product = &Product{ID: partialData.ID}
				bufer[partialData.ID] = product
			}
			product.Price = partialData.Price
			receivedPartsCount[partialData.ID]++

		case partialData, ok := <-descChan:
			if !ok {
				closedChannels++
				descChan = nil
				continue
			}
			product, exists := bufer[partialData.ID]
			if !exists {
				product = &Product{ID: partialData.ID}
				bufer[partialData.ID] = product
			}
			product.Description = partialData.Description
			receivedPartsCount[partialData.ID]++

		case partialData, ok := <-availChan:
			if !ok {
				closedChannels++
				availChan = nil
				continue
			}
			product, exists := bufer[partialData.ID]
			if !exists {
				product = &Product{ID: partialData.ID}
				bufer[partialData.ID] = product
			}
			product.Availability = partialData.Availability
			receivedPartsCount[partialData.ID]++
		}

		for id, count := range receivedPartsCount {
			if count == channelsTotal {
				finalProducts <- *bufer[id]
				delete(bufer, id)
				delete(receivedPartsCount, id)
			}
		}

		if closedChannels == channelsTotal && len(bufer) == 0 {
			break
		}
	}

	close(finalProducts)

}
