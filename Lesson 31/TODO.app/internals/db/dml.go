package db

const (
	SetTaskPrioritiesTableValuesQuery = `INSERT INTO task_priorities (priority)
										 VALUES ('Urgent'),
												('High'),
												('Medium'),
												('Low'),
												('Backlog');`
	CreateNewUserQuery = `INSERT INTO users (name)
						  VALUES ($1);`
	FindUserByNameQuery = `SELECT id, name FROM users WHERE name ILIKE ('%$1%');`
	UpdateUserByIDQuery = `UPDATE users SET name = $1 WHERE id = $2;`
	DeleteUserByIDQuery = `DELETE FROM users WHERE id = $1;`
	CreateTaskQuery     = `INSERT INTO tasks (title, description, task_priority_id, user_id)
						   VALUES ($1,$2,$3,$4);`
	GetAllTasksQuery = `SELECT t.id, 
							   t.is_deleted, 
							   t.description, 
							   t.task_priority_id,
							   tp.priority,
							   t.user_id,
							   u.name,
							   t.created_at, 
							   t.is_done, 
							   t.done_at, 
							   t.is_deleted, 
							   t.deleted_at
						FROM tasks t
								 JOIN task_priorities tp ON t.task_priority_id = tp.id
								 JOIN users u ON t.user_id = u.id;`
	GetTasksByIDQuery = `SELECT t.id, 
							   t.is_deleted, 
							   t.description, 
							   t.task_priority_id,
							   tp.priority,
							   t.user_id,
							   u.name,
							   t.created_at, 
							   t.is_done, 
							   t.done_at, 
							   t.is_deleted, 
							   t.deleted_at
						FROM tasks t
								 JOIN task_priorities tp ON t.task_priority_id = tp.id
								 JOIN users u ON t.user_id = u.id
						WHERE t.id = $1;`
)
