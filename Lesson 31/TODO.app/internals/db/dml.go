package db

const (
	CreateTaskQuery = `INSERT INTO tasks (title, description, task_priority_id, user_name)
						   VALUES ($1,$2,$3,$4);`
	GetAllTasksQuery = `SELECT t.id, 
							   t.title, 
							   t.description, 
							   tp.priority,
							   t.user_name,
							   t.created_at, 
							   t.is_done, 
							   t.is_deleted
						FROM tasks t
								 JOIN task_priorities tp ON t.task_priority_id = tp.id`
	GetTasksByIDQuery = `SELECT t.id, 
							   t.title, 
							   t.description, 
							   tp.priority,
							   t.user_name,
							   t.created_at, 
							   t.is_done, 
							   t.is_deleted
						FROM tasks t
								 JOIN task_priorities tp ON t.task_priority_id = tp.id
						WHERE t.id = $1;`
	DoneTaskByIDQuery = `UPDATE tasks
							SET is_done = TRUE, done_at = CURRENT_TIMESTAMP
							WHERE id = $1;`
	UndoneTaskByIDQuery = `UPDATE tasks
							SET is_done = false, done_at = NULL
							WHERE id = $1;`
	DeleteTaskByIDQuery = `UPDATE tasks
							SET is_deleted = TRUE, deleted_at = CURRENT_TIMESTAMP
							WHERE id = $1;`
	UpdateTaskTitleByIDQuery = `UPDATE tasks
										SET title = $1
										WHERE id = $2;`
	UpdateTaskDescriptionByIDQuery = `UPDATE tasks
										SET description = $1
										WHERE id = $2;`
	CancelTaskDeletionByIDQuery = `UPDATE tasks
							SET is_deleted = FALSE, deleted_at = NULL
							WHERE id = $1;`
	GetDoneTasksQuery = `SELECT t.id, 
							   t.title, 
							   t.description, 
							   tp.priority,
							   t.user_name,
							   t.created_at, 
							   t.is_done, 
							   t.done_at, 
							   t.is_deleted, 
						FROM tasks t
								 JOIN task_priorities tp ON t.task_priority_id = tp.id
								 WHERE t.is_done = TRUE;`
	GetDeletedTasksQuery = `SELECT t.id, 
							   t.title, 
							   t.description, 
							   tp.priority,
							   t.user_name,
							   t.created_at, 
							   t.is_done, 
							   t.done_at, 
							   t.is_deleted, 
							   t.deleted_at
						FROM tasks t
								 JOIN task_priorities tp ON t.task_priority_id = tp.id
								 WHERE t.deleted = TRUE;`
	UpdateTaskUserNameByIDQuery = `UPDATE tasks
									SET user_name = $1
									WHERE id = $2;`
	UpdateTaskPriorityByIDQuery = `UPDATE tasks
										SET priority_id = $1
										WHERE id = $2;`
)
