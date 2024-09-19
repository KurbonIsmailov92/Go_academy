package main

import (
	"todo-crud/internals/db"
	"todo-crud/logger"
	"todo-crud/pkg/controllers"
)

func main() {
	if err := logger.Init(); err != nil {
		panic(err)
	}
	if err := db.ConnectToDB(); err != nil {
		panic(err)
	}
	if err := db.MigrateTables(); err != nil {
		panic(err)
	}
	if err := db.InsertSeeds(); err != nil {
		panic(err)
	}
	defer func() {
		err := db.CloseDBConnection()
		if err != nil {
		}
	}()
	if err := controllers.RunRoutes(); err != nil {
		panic(err)
	}
}
