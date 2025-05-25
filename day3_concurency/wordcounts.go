package day3concurency

import (
	"strings"
	"sync"
)

func CountWords(text string, wordCountChan chan map[string]int, wg *sync.WaitGroup) {

	words := strings.Fields(text)
	counter := make(map[string]int)
	for _, val := range words {
		counter[strings.ToLower(val)]++
	}

	wordCountChan <- counter
}
