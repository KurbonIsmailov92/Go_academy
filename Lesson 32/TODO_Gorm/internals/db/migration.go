package db

import (
	"todo-gorm/internals/models"
)

func MigrateTables() error {
	err := dbConnection.AutoMigrate(models.Task{}, models.TaskPriority{})
	if err != nil {
		return err
	}
	return nil
}
