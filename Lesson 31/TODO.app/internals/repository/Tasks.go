package repository

import (
	"todo_app/internals/db"
	"todo_app/internals/models"
)

func SetNewTaskToDB(task models.Task) error {
	_, err := db.GetDBConnection().
		Exec(db.CreateTaskQuery, task.Title, task.Description, task.TaskPriorityID, task.UserID)
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
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.TaskPriorityID, &task.TaskPriority,
			&task.UserID, &task.UserName, &task.CreatedAt, &task.IsDone, &task.DoneAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskByID(taskID int) (task models.Task, err error) {

	row := db.GetDBConnection().QueryRow(db.GetTasksByIDQuery, taskID)
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.TaskPriorityID, &task.TaskPriority,
		&task.UserID, &task.UserName, &task.CreatedAt, &task.IsDone, &task.DoneAt)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}
