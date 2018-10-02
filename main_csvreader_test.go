package main

import (
	"fmt"
	//	"log"
	"testing"
)

func TestReadCSVFileByLine(t *testing.T) {
	ReadCSVFileByLine("./sampledata/5kSalesRecords.csv", false, defaultReturn)
}

func defaultReturn(m string) {
	fmt.Println(m)
}
