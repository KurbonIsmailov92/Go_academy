package repository

import (
	"todo_app/internals/db"
	"todo_app/internals/models"
)

func SetNewTaskToDB(task models.Task) error {
	_, err := db.GetDBConnection().
		Exec(db.CreateTaskQuery, task.Title, task.Description, task.TaskPriorityID, task.UserName)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	rows, err := db.GetDBConnection().Query(db.GetAllTasksQuery)
	if err != nil {
		return tasks, err
	}
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.TaskPriority,
			&task.UserName, &task.CreatedAt, &task.IsDone, &task.DoneAt, &task.IsDeleted, &task.DeletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskByID(taskID int) (task models.Task, err error) {

	row := db.GetDBConnection().QueryRow(db.GetTasksByIDQuery, taskID)
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.TaskPriority,
		&task.UserName, &task.CreatedAt, &task.IsDone, &task.DoneAt, &task.IsDeleted, &task.DeletedAt)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func UpdateTaskTitleByID(taskID int, newTitle string) error {
	_, err := db.GetDBConnection().Exec(db.UpdateTaskTitleByIDQuery, newTitle, taskID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTaskUserByID(taskID int, newUsername string) error {
	_, err := db.GetDBConnection().Exec(db.UpdateTaskUserNameByIDQuery, newUsername, taskID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTaskPriorityByID(taskID int, newPrID int) error {
	_, err := db.GetDBConnection().Exec(db.UpdateTaskPriorityByIDQuery, newPrID, taskID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTaskDescriptionByID(taskID int, newDescription string) error {
	_, err := db.GetDBConnection().Exec(db.UpdateTaskDescriptionByIDQuery, newDescription, taskID)
	if err != nil {
		return err
	}
	return nil
}

func DoneTaskByID(taskID int) error {
	_, err := db.GetDBConnection().Exec(db.DoneTaskByIDQuery, taskID)
	if err != nil {
		return err
	}
	return nil

}

func UnDoneTaskByID(taskID int) error {
	_, err := db.GetDBConnection().Exec(db.UndoneTaskByIDQuery, taskID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTaskByID(taskID int) error {
	_, err := db.GetDBConnection().Exec(db.DeleteTaskByIDQuery, taskID)
	if err != nil {
		return err
	}
	return nil
}

func CancelTaskDeletionByID(taskID int) error {
	_, err := db.GetDBConnection().Exec(db.CancelTaskDeletionByIDQuery, taskID)
	if err != nil {
		return err
	}
	return nil
}

func GetDoneTasks() ([]models.Task, error) {
	var tasks []models.Task
	rows, err := db.GetDBConnection().Query(db.GetDoneTasksQuery)
	if err != nil {
		return tasks, err
	}
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.TaskPriority,
			&task.UserName, &task.CreatedAt, &task.IsDone, &task.DoneAt, &task.IsDeleted, &task.DeletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetDeletedTasks() ([]models.Task, error) {
	var tasks []models.Task
	rows, err := db.GetDBConnection().Query(db.GetDeletedTasksQuery)
	if err != nil {
		return tasks, err
	}
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.TaskPriority,
			&task.UserName, &task.CreatedAt, &task.IsDone, &task.DoneAt, &task.IsDeleted, &task.DeletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
