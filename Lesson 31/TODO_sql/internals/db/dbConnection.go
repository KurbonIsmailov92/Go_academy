package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var dbConnection *sql.DB

func ConnectToDB() error {
	connectionString := "user=kurbon password=ismoil dbname=todo_db sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	dbConnection = db
	return nil
}

func CloseDBConnection() error {
	err := dbConnection.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetDBConnection() *sql.DB {
	return dbConnection
}
