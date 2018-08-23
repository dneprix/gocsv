package parser

import (
	"encoding/csv"
	"os"
)

// Csv parses csv file by filepath and returns records
func Csv(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// CsvPrint get records and print to STDOUT as csv file
func CsvPrint(records [][]string) {
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)
}
