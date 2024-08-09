package db

import "todo-gorm/internals/models"

func InsertSeeds() error {
	var taskPriorities []models.TaskPriority

	taskPriorities = append(taskPriorities, models.TaskPriority{Priority: "Urgent"})
	taskPriorities = append(taskPriorities, models.TaskPriority{Priority: "High"})
	taskPriorities = append(taskPriorities, models.TaskPriority{Priority: "Medium"})
	taskPriorities = append(taskPriorities, models.TaskPriority{Priority: "Low"})
	taskPriorities = append(taskPriorities, models.TaskPriority{Priority: "Backlog"})

	if err := dbConnection.Create(&taskPriorities).Error; err != nil {
		return err
	}

	return nil
}
