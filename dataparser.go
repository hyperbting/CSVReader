package main

import (
	"strconv"
	"strings"
)

//http://eforexcel.com/wp/downloads-18-sample-csv-files-data-sets-for-testing-sales/
type SampleData struct {
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

//SampleDataParser parse string to sampleData
func SampleDataParser(msg string) SampleData {
	vals := strings.Split(msg, ",")

	//TODO: Valid value here
	oid, _ := strconv.ParseInt(vals[5], 10, 64)
	uSold, _ := strconv.ParseInt(vals[7], 10, 64)

	uPrice, _ := strconv.ParseFloat(vals[8], 64)
	uCost, _ := strconv.ParseFloat(vals[9], 64)

	tRevenue, _ := strconv.ParseFloat(vals[10], 64)
	tCost, _ := strconv.ParseFloat(vals[11], 64)
	tProfit, _ := strconv.ParseFloat(vals[12], 64)

	return SampleData{
		Country:       vals[0],
		ItemType:      vals[1],
		SalesChannel:  vals[2],
		OrderPriority: vals[3],
		OrderDate:     vals[4],
		OrderID:       int(oid),
		ShipDate:      vals[6],
		UnitsSold:     int(uSold),
		UnitPrice:     uPrice,
		UnitCost:      uCost,
		TotalRevenue:  tRevenue,
		TotalCost:     tCost,
		TotalProfit:   tProfit}
}
