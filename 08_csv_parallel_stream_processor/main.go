package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Record struct {
	id            string
	name          string
	ISIN          string
	units         int
	purchase_date string
}

// processRecord simulates processing of a record
func processRecord(r []string) Record {
	fmt.Printf("Processing record ID: %s\n", r[0])

	var record Record

	units, errors := strconv.Atoi(strings.ReplaceAll(r[3], " ", ""))
	if errors != nil {
		fmt.Println("Error:", errors)
	} else {
		record = Record{r[0], r[1], r[2], units, r[4]}
	}

	// Simulate some work with sleep
	time.Sleep(time.Millisecond * 1000)

	return record
}

func main() {
	file, err := os.Open("your_file.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new CSV reader from the file
	reader := csv.NewReader(file)

	// Skip the headers
	if _, err := reader.Read(); err != nil {
		log.Fatalf("Failed to read header: %s", err)
	}

	var wg sync.WaitGroup
	var recordsArray []Record
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to read record: %s", err)
		}

		wg.Add(1) // Increment the WaitGroup counter
		go func(r []string) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			res := processRecord(r)

			recordsArray = append(recordsArray, res)
		}(record) // Pass the current record to the goroutine
	}

	wg.Wait()
	fmt.Println(recordsArray)
}
