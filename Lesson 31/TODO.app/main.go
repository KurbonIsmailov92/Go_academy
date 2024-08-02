package main

import "todo_app/internals/db"

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

}
