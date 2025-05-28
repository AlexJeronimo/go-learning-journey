package day3concurency

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func GenerateLogEntries(filename string, count int) error {

	levels := []string{"INFO", "WARN", "ERROR"}
	messages := []string{
		"User logged in successfully",
		"File not found in specified directory",
		"Database connection error: Could not connect to primary replica",
		"API call received with invalid parameters",
		"Service restarted due to unhandled exception",
		"Cache invalidated for key 'user_data'",
		"Processing batch of 100 items",
		"Authentication failed for user 'guest'",
		"Resource limit exceeded for process",
		"Configuration reloaded successfully",
	}

	var logLines []string

	fmt.Printf("Log generator: start generating %d rows to file %s...\n", count, filename)
	for i := 0; i < count; i++ {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		logLevel := levels[rand.Intn(len(levels))]
		requestID := rand.Intn(9000) + 1000
		logMessage := messages[rand.Intn(len(messages))]

		logEntry := fmt.Sprintf("%s [%s] - %s (RequestID: %04d)", currentTime, logLevel, logMessage, requestID)

		logLines = append(logLines, logEntry)
	}

	finalLogContent := strings.Join(logLines, "\n")

	logData := []byte(finalLogContent)

	err := os.WriteFile(filename, logData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadLogEntries(filename string, lines chan<- string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}

	close(lines)
}

func LogAnalyzer(id int, in <-chan string, erros chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range in {
		time.Sleep(5 * time.Millisecond)
		if strings.Contains(data, "[ERROR]") {
			erros <- data
		}

		fmt.Printf("Worker %d: completed row <%s>...\n", id, data[:30])
	}

}
