package main

import (
	"encoding/csv"
	"fmt"
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
func processRecord(r Record) {
	fmt.Printf("Processing record ID: %s\n", r.id)
	// Simulate some work with sleep
	time.Sleep(time.Millisecond * 1000)
}

func readFile() [][]string {
	file, err := os.Open("your_file.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new CSV reader from the file
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	return records
}

func main() {

	// Read all records from the CSV file
	records := readFile()

	var recordsArray []Record
	// Print each record
	for index, record := range records {

		if index == 0 {
			continue
		}

		units, errors := strconv.Atoi(strings.ReplaceAll(record[3], " ", ""))
		if errors != nil {
			fmt.Println("Error:", errors)
		} else {
			recordsArray = append(recordsArray, Record{record[0], record[1], record[2], units, record[4]})
		}
	}

	var wg sync.WaitGroup

	for _, record := range recordsArray {
		wg.Add(1) // Increment the WaitGroup counter
		go func(r Record) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			processRecord(r)
		}(record) // Pass the current record to the goroutine
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All records processed.")
}
