package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func ConnectToDB() error {
	connectionString := "host=localhost port=5432 user=kurbon dbname=todo_db password=ismoil"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	dbConnection = db
	return nil
}

func CloseDBConnection() error {

	return nil
}

func GetDBConnection() *gorm.DB {
	return dbConnection
}
