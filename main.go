package main

import (
	//"bufio"
	"fmt"
	"log"
	"time"
)

type rect struct {
	width, height int
}

func main() {

	config := LoadConfig()

	// read from console for the csv files
	fmt.Print("Enter file path: [./5kSalesRecords.csv]")
	var filPath string
	fmt.Scanln(&filPath)
	if len(filPath) <= 0 {
		filPath = "./5kSalesRecords.csv"
	}

	// // Start a worker goroutine, giving it the channel to notify on.
	formatedRow := make(chan SampleData, 50) // Chan
	readFinished := make(chan bool, 1)       // Chan
	processFinished := make(chan bool, 1)    // Chan
	sqlActorFinished := make(chan bool, 1)   // Chan

	rowStoragePtr := new([]SampleData) // Ptr

	log.Println("formatedRow Chan prepared")

	config.CSVReaderID = 100
	config.FinishedChan = readFinished
	go csvReadWorker(config, filPath, formatedRow)

	config.ProcessWorkerID = 300
	config.FinishedChan = processFinished
	go csvProcessWorker(config, formatedRow, rowStoragePtr)

	config.SQLInserterID = 500
	config.FinishedChan = sqlActorFinished
	go SQLInserionWorker(config, rowStoragePtr)

	log.Println("Workers ready...")

	<-readFinished
	log.Println("wait for ProcessWorker to Finish")
	<-processFinished
	log.Println("wait for SQLWorker to Finish")
	<-sqlActorFinished

	log.Println("done... Press Any Key to Leave")
	fmt.Scanln()

}

func csvReadWorker(config Configuration, filePath string, formatedRow chan<- SampleData) {
	log.Println("csvReadWorker", "r"+string(config.CSVReaderID), "reading...")

	ReadCSVFileByLine(filePath, true, func(msg string) { formatedRow <- SampleDataParser(msg) })

	log.Println("csvReadWorker", config.CSVReaderID, "offline")

	config.FinishedChan <- true
	close(formatedRow)
}

func csvProcessWorker(config Configuration, singleRow <-chan SampleData, rowStorage *[]SampleData) {
	log.Println("csvProcessWorker", config.ProcessWorkerID, "propressing...")

	for {
		if len(*rowStorage) >= 5000 {
			time.Sleep(time.Second * 1)
		}

		// more just indicate the chan is not closed
		j, more := <-singleRow
		if more {
			// put into tmp storage
			*rowStorage = append(*rowStorage, j)
		} else {
			log.Println("csvProcessWorker", config.ProcessWorkerID, "offline")
			config.FinishedChan <- true
			return
		}
	}
}

// SQLInserionWorker 123
func SQLInserionWorker(config Configuration, ptr *[]SampleData) {

	idleCounter := 0

	for {
		time.Sleep(time.Second * 1)

		if idleCounter > 5 {
			break // break for
		}

		storedRows := len(*ptr)
		if storedRows <= 0 {
			time.Sleep(time.Second)
			idleCounter++
			continue
		}

		idleCounter = 0
		var tmp []SampleData

		if storedRows > config.MaxSQLInsertionRow {
			tmp = (*ptr)[:config.MaxSQLInsertionRow]
			*ptr = (*ptr)[config.MaxSQLInsertionRow:storedRows]
		} else {
			tmp = *ptr
			*ptr = (*ptr)[:0]
		}

		// send tmp for SQL insertion
		fmt.Println("SQLInserionWorker SQL", len(tmp), "lines")
	}

	log.Println("SQLInserionWorker", config.SQLInserterID, "offline")
	config.FinishedChan <- true
}
