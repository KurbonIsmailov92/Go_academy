package service

import (
	"todo-crud/internals/models"
	"todo-crud/pkg/repository"
)

func CreateNewTask(task models.Task) error {
	if err := repository.SetNewTaskToDB(task); err != nil {
		return err
	}
	return nil
}

func GetAllTasks() (task []models.Task, err error) {
	task, err = repository.GetAllTasks()
	if err != nil {
		return task, err
	}
	return task, nil
}

func GetTaskByID(taskId int) (task models.Task, err error) {
	task, err = repository.GetTaskByID(taskId)
	if err != nil {
		return task, err
	}
	return task, nil
}

func UpdateTaskByID(taskID int, task models.Task) error {
	existTask, err := repository.GetTaskByID(taskID)
	if err != nil {
		return err
	}
	task.ID = existTask.ID
	if err = repository.UpdateTaskByID(taskID, task, existTask); err != nil {
		return err
	}
	return nil
}

func DoneTaskByID(taskID int) error {
	if err := repository.DoneTaskByID(taskID); err != nil {
		return err
	}
	return nil
}

func UnDoneTaskByID(taskID int) error {
	if err := repository.UnDoneTaskByID(taskID); err != nil {
		return err
	}
	return nil
}

func DeleteTaskByID(taskID int) error {
	if err := repository.DeleteTaskByID(taskID); err != nil {
		return err
	}
	return nil
}

func CancelTaskDeletionByID(taskID int) error {
	if err := repository.CancelTaskDeletionByID(taskID); err != nil {
		return err
	}
	return nil
}

func GetDoneTasks() ([]models.Task, error) {
	tasks, err := repository.GetDoneTasks()
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func GetDeletedTasks() ([]models.Task, error) {
	tasks, err := repository.GetDeletedTasks()
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}
