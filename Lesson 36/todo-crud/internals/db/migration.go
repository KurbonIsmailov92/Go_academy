package db

import (
	models2 "todo-crud/internals/models"
)

func MigrateTables() error {
	err := dbConnection.AutoMigrate(models2.Task{}, models2.TaskPriority{}, models2.User{})
	if err != nil {
		return err
	}
	return nil
}
