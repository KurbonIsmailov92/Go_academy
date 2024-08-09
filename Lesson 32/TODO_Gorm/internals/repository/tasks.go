package repository

import (
	"time"
	"todo-gorm/internals/db"
	"todo-gorm/internals/models"
)

func SetNewTaskToDB(t models.Task) error {
	err := db.GetDBConnection().Create(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllTasks() (tasks []models.Task, err error) {
	tasks = []models.Task{}
	err = db.GetDBConnection().Preload("TaskPriority").Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskByID(taskID int) (task models.Task, err error) {

	err = db.GetDBConnection().Raw(db.GetTasksByIDQuery, taskID).Scan(&task).Error
	if err != nil {
		return models.Task{}, err
	}

	return task, nil

}

func UpdateTaskTitleByID(taskID int, newTitle string) error {
	db.GetDBConnection().Model(&models.Task{}).Where("id = ?", taskID).Update("title", newTitle)
	return nil
}

func UpdateTaskUserByID(taskID int, newUsername string) error {
	db.GetDBConnection().Model(&models.Task{}).Where("id = ?", taskID).Update("user_name", newUsername)
	return nil
}

func UpdateTaskPriorityByID(taskID int, newPrID int) error {
	db.GetDBConnection().Model(&models.Task{}).Where("id = ?", taskID).Update("priority_id", newPrID)
	return nil
}

func UpdateTaskDescriptionByID(taskID int, newDescription string) error {
	db.GetDBConnection().Model(&models.Task{}).Where("id = ?", taskID).Update("description", newDescription)
	return nil
}

func DoneTaskByID(taskID int) error {
	db.GetDBConnection().Table("tasks").Where("id = ?", taskID).Update("is_done", true)
	db.GetDBConnection().Table("tasks").Where("id = ?", taskID).Update("done_at", time.Now())
	return nil

}

func UnDoneTaskByID(taskID int) error {
	db.GetDBConnection().Table("tasks").Where("id = ?", taskID).Update("is_done", false)
	db.GetDBConnection().Table("tasks").Where("id = ?", taskID).Update("done_at", nil)
	return nil

}

func DeleteTaskByID(taskID int) error {
	db.GetDBConnection().Table("tasks").Where("id = ?", taskID).Update("is_deleted", true)
	db.GetDBConnection().Table("tasks").Where("id = ?", taskID).Update("deleted_at", time.Now())
	return nil
}

func CancelTaskDeletionByID(taskID int) error {
	db.GetDBConnection().Table("tasks").Where("id = ?", taskID).Update("is_deleted", false)
	db.GetDBConnection().Table("tasks").Where("id = ?", taskID).Update("deleted_at", nil)
	return nil

}

func GetDoneTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.GetDBConnection().Raw(db.GetDoneTasksQuery).Scan(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func GetDeletedTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.GetDBConnection().Raw(db.GetDeletedTasksQuery).Scan(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}
