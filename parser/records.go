package parser

import (
	"strconv"
	"strings"

	"github.com/dneprix/gocsv/calc"
)

const abc = "abcdefghijklmnopqrstuvwxyz"
const errorValue = "#ERR"

// Records function gets values from expressions for all records
func Records(records [][]string) [][]string {
	for row, cols := range records {
		for col, record := range cols {
			records[row][col] = errorValue // fix for circle refs
			records[row][col] = recordValue(record, records)
		}
	}
	return records
}

// Get value for record
func recordValue(record string, records [][]string) string {
	recordWithRefs := recordRefs(record, records)
	if r := calc.Rpn(recordWithRefs); len(r) > 0 {
		return r
	}
	return errorValue
}

// Replace refs to values in record expression
func recordRefs(record string, records [][]string) string {
	items := strings.Fields(strings.ToLower(record))
	for i, itm := range items {
		firstSymbol := itm[0]
		if firstSymbol >= 'a' && firstSymbol <= 'z' {
			colIndex := strings.Index(abc, string(firstSymbol))
			rowIndex, err := strconv.Atoi(itm[1:])
			if err != nil || rowIndex > len(records) || colIndex+1 > len(records[rowIndex-1]) {
				return errorValue
			}

			rec := strings.TrimSpace(records[rowIndex-1][colIndex])
			if _, err := strconv.ParseFloat(rec, 64); err == nil {
				items[i] = rec
			} else {
				items[i] = recordValue(rec, records)
			}
		}
	}
	return strings.Join(items, " ")
}
