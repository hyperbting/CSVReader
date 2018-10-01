package main

import (
	"bufio"
	"log"
	"os"
)

//http://eforexcel.com/wp/downloads-18-sample-csv-files-data-sets-for-testing-sales/
type sampleData struct {
	Region        string
	Country       string
	ItemType      string
	SalesChannel  string
	OrderPriority string
	OrderDate     string
	OrderID       int
	ShipDate      string
	UnitsSold     int
	UnitPrice     float64
	UnitCost      float64
	TotalRevenue  float64
	TotalCost     float64
	TotalProfit   float64
}

//ReadCSVFileByLine Read CSV data
func ReadCSVFileByLine(filePath string, ignoreFirstLine bool, rowUser func(string)) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	scanner := bufio.NewScanner(csvFile)
	for scanner.Scan() {

		if ignoreFirstLine {
			ignoreFirstLine = false
			continue
		}

		rowUser(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
