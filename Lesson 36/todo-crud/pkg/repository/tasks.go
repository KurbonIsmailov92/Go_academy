package repository

import (
	"time"
	"todo-crud/internals/db"
	"todo-crud/internals/models"
	"todo-crud/logger"
)

func SetNewTaskToDB(t models.Task) error {
	if err := db.GetDBConnection().Create(&t).Error; err != nil {
		logger.Error.Printf("[repository] Error making new task: %v", err)
		return err
	}

	return nil
}

func GetAllTasks() (tasks []models.Task, err error) {
	tasks = []models.Task{}
	if err = db.GetDBConnection().Preload("TaskPriority").
		Preload("User").
		Order("id").
		Find(&tasks).Error; err != nil {
		logger.Error.Printf("[repository] Error getting all tasks: %v", err)
		return nil, err
	}

	return tasks, nil
}

func GetTaskByID(taskID int) (task models.Task, err error) {

	if err = db.GetDBConnection().Preload("TaskPriority").
		Preload("User").
		Where("id = ?", taskID).
		First(&task).Error; err != nil {
		logger.Error.Printf("[repository] Error getting task: %v", err)
		return models.Task{}, err
	}

	return task, nil

}

func UpdateTaskByID(taskID int, task, existTask models.Task) error {
	if err := db.GetDBConnection().Model(&existTask).
		Where("id = ?", taskID).
		Updates(task).Error; err != nil {
		logger.Error.Printf("[repository] Error updating task: %v", err)
		return err
	}

	return nil
}

func DoneTaskByID(taskID int) error {
	if err := db.GetDBConnection().Model(models.Task{}).
		Where("id", taskID).
		Updates(map[string]interface{}{
			"is_done": true,
			"done_at": time.Now(),
		}).Error; err != nil {
		logger.Error.Printf("[repository] Error updating task: %v", err)
		return err
	}

	return nil
}

func UnDoneTaskByID(taskID int) error {
	if err := db.GetDBConnection().Model(models.Task{}).
		Where("id", taskID).
		Updates(map[string]interface{}{
			"is_done": false,
			"done_at": nil,
		}).Error; err != nil {
		logger.Error.Printf("[repository] Error updating task: %v", err)
		return err
	}

	return nil
}

func DeleteTaskByID(taskID int) error {
	if err := db.GetDBConnection().Model(models.Task{}).
		Where("id", taskID).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": time.Now(),
		}).Error; err != nil {
		logger.Error.Printf("[repository] Error updating task: %v", err)
		return err
	}

	return nil
}

func CancelTaskDeletionByID(taskID int) error {
	if err := db.GetDBConnection().Model(models.Task{}).
		Where("id", taskID).
		Updates(map[string]interface{}{
			"is_deleted": false,
			"deleted_at": nil,
		}).Error; err != nil {
		logger.Error.Printf("[repository] Error returning task: %v", err)
		return err
	}

	return nil
}

func GetDoneTasks() ([]models.Task, error) {
	tasks := []models.Task{}
	if err := db.GetDBConnection().Preload("TaskPriority").
		Preload("User").
		Where("is_done", true).
		Find(&tasks).Error; err != nil {
		logger.Error.Printf("[repository] Error getting done tasks: %v", err)
		return nil, err
	}

	return tasks, nil
}

func GetDeletedTasks() ([]models.Task, error) {
	tasks := []models.Task{}
	if err := db.GetDBConnection().Preload("TaskPriority").
		Preload("User").
		Where("is_deleted", true).
		Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
