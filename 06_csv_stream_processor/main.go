package main

import (
	"encoding/csv"
	"fmt"
	"io"
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

	// Skip the headers
	if _, err := reader.Read(); err != nil {
		log.Fatalf("Failed to read header: %s", err)
	}

	var recordsArray []Record
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to read record: %s", err)
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
