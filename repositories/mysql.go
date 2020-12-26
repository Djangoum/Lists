package repositories

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func CreateDatabaseConnection() *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", "root", "rootpw", "localhost", "lists")
	localDb, err := sql.Open("mysql", dataSourceName)

	DB = localDb

	if err != nil {
		log.Fatal(err.Error())
	}

	return DB
}
