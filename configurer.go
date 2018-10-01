package main

import (
	"github.com/tkanos/gonfig"
	"log"
)

// Configuration is a file saved locally
type Configuration struct {
	CSVReaderID     int
	ProcessWorkerID int
	SQLInserterID   int

	FinishedChan chan<- bool

	MaxSQLInsertionRow int `json:"maxRowCount"`
}

// LoadConfig load configure file from HDD
func LoadConfig() (config Configuration) {

	// read config file
	err := gonfig.GetConf("config.json", &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

// LoadMySQLConfig load configure file from HDD
func LoadMySQLConfig() (config DBConnectionParameter) {

	// read config file
	err := gonfig.GetConf("mysqlconfig.json", &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
