package main

import (
	"todo-gorm/appRun"
	"todo-gorm/internals/db"
)

func main() {
	err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := db.CloseDBConnection()
		if err != nil {
			panic(err)
		}
	}()

	err = db.MigrateTables()
	if err != nil {
		panic(err)
	}

	err = db.InsertSeeds()
	if err != nil {
		panic(err)
	}

	appRun.Run()

}
