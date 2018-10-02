package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// DBConnectionParameter ...
type DBConnectionParameter struct {
	Username string `json:"mysqlUserName"`
	Password string `json:"mysqlPassword"`
	DBURL    string `json:"mysqlURL"`
	DBPort   string `json:"mysqlPort"`
	DBName   string `json:"mysqlDBNAme"`
}

var db *sql.DB
var dbParameter DBConnectionParameter

// Setup General Init of DB parameter
func Setup(dbuname string, dbpass string, targetURL string, targetPort string, dbtable string) {
	dbParameter.Username = dbuname
	dbParameter.Password = dbpass
	dbParameter.DBURL = targetURL
	dbParameter.DBPort = targetPort
	dbParameter.DBName = dbtable
}

// SetupByStruct General Init of DB parameter
func SetupByStruct(config DBConnectionParameter) {
	dbParameter.Username = config.Username
	dbParameter.Password = config.Password
	dbParameter.DBURL = config.DBURL
	dbParameter.DBPort = config.DBPort
	dbParameter.DBName = config.DBName
}

// InitDB General Init of DB connection
func InitDB() {

	var err error
	db, err := sql.Open("mysql", dbParameter.Username+":"+dbParameter.Password+"@tcp("+dbParameter.DBURL+":"+dbParameter.DBPort+")/"+dbParameter.DBName)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

// SetupNInitDB load local config and steup
func SetupNInitDB() {
	config := LoadMySQLConfig()
	SetupByStruct(config)
	InitDB()
}

// CloseDB close DB connection, not used in general
func CloseDB() {
	defer db.Close()
}

// GetDB Direct usage og db connection
func GetDB() *sql.DB {
	if db == nil {
		SetupNInitDB()
	}

	return db
}

// CreateDeviceUser create an account with GAccount OR AppleID
func CreateDeviceUser(datas []SampleData) (res bool, dbErr error) {

	stmtString := "INSERT INTO users (name, %s, fasecret, status, updated_at) VALUES (?, ?, ?, ?, ?)"
	stmt, dbErr := GetDB().Prepare(stmtString)
	CloseDB()

	if dbErr != nil {
		log.Fatal(dbErr)
		return res, dbErr
	}
	defer stmt.Close()

	// sqlRes, dbErr := stmt.Exec(uid.GetName(), uid.GetID(), res.FASecret, "active", time.Now())
	// if dbErr != nil {
	// 	log.Fatal(dbErr)
	// }

	// LastInsertId, dbErr := sqlRes.LastInsertId()
	// if dbErr != nil {
	// 	log.Fatal(dbErr)
	// }

	return res, dbErr
}
