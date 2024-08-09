package db

import (
	"errors"
)

func MigrateTables() error {

	tx, err := dbConnection.Begin()
	if err != nil {
		return errors.New("Failed to begin transaction: " + err.Error())
	}

	_, err = tx.Exec(CreateTaskPrioritiesTableQuery)
	if err != nil {
		tx.Rollback()
		return errors.New("Failed to create Priorities table: " + err.Error())
	}

	_, err = tx.Exec(CreateTasksTableQuery)
	if err != nil {
		tx.Rollback()
		return errors.New("Failed to create tasks table: " + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("Failed to commit transaction: " + err.Error())
	}
	return nil
}
