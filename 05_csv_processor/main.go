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

func main() {
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

	var recordsArray []Record
	// Print each record
	for index, record := range records {

		if index == 0 {
			continue
		}

		units, errors := strconv.Atoi(strings.ReplaceAll(record[3], " ", ""))
		if errors != nil {
			fmt.Println("Error:", err)
		} else {
			recordsArray = append(recordsArray, Record{record[0], record[1], record[2], units, record[4]})
		}
	}

	fmt.Println(recordsArray)
}
