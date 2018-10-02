package main

import (
	"bufio"
	"log"
	"os"
)

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
