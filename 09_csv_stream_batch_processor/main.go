package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	return record
}

func processBatch(batch [][]string) []Record {
	var result []Record
	for _, record := range batch {
		result = append(result, processRecord(record))
	}

	return result
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

	batchSize := 10
	batch := make([][]string, 0, batchSize)

	var recordsArray []Record
	for {
		record, err := reader.Read()

		if err != nil {
			if err.Error() == "EOF" {
				// Process the remaining records in the last batch
				if len(batch) > 0 {
					recordsArray = append(recordsArray, processBatch(batch)...)
				}
				break
			}
			log.Fatalf("Failed to read line: %s", err)
		}

		if err != nil {
			log.Fatalf("Failed to read record: %s", err)
		}

		// Append the record to the current batch
		batch = append(batch, record)

		// If the batch is full, process it and reset the batch
		if len(batch) == batchSize {
			recordsArray = append(recordsArray, processBatch(batch)...)
			batch = batch[:0] // Reset the batch to an empty slice
		}

	}

	fmt.Println(recordsArray)
}
