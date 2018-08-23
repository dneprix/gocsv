package main

import (
	"fmt"
	"os"

	"github.com/dneprix/gocsv/parser"
)

func main() {
	// Get arg with file path
	if len(os.Args) < 2 {
		fmt.Println("ERROR:", "No file")
		return
	}
	filepath := os.Args[1]

	// Get records from csv file
	records, err := parser.Csv(filepath)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Parse records and get result with values
	result := parser.Records(records)

	// Print result CSV to STDOUT
	parser.CsvPrint(result)
}
