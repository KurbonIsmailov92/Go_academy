package db

const (
	GetTasksByIDQuery = `SELECT t.id, 
							   t.title, 
							   t.description, 
							   tp.priority,
							   t.user_name,
							   t.created_at, 
							   t.is_done,
							   COALESCE(t.done_at, '0001-01-01 00:00:00') AS done_at,
							   t.is_deleted,
							   COALESCE(t.deleted_at, '0001-01-01 00:00:00') AS deleted_at
						FROM tasks t
								 JOIN task_priorities tp ON t.priority_id = tp.id
						WHERE t.id = $1;`
	GetDoneTasksQuery = `SELECT t.id, 
							   t.title, 
							   t.description, 
							   tp.priority,
							   t.user_name,
							   t.created_at, 
							   t.is_done,
							   COALESCE(t.done_at, '0001-01-01 00:00:00') AS done_at,
							   t.is_deleted,
							   COALESCE(t.deleted_at, '0001-01-01 00:00:00') AS deleted_at
						FROM tasks t
								 JOIN task_priorities tp ON t.priority_id = tp.id
								 WHERE t.is_done = TRUE;`
	GetDeletedTasksQuery = `SELECT t.id, 
								   t.title, 
								   t.description, 
								   tp.priority,
								   t.user_name,
								   t.created_at, 
								   t.is_done,
								   COALESCE(t.done_at, '0001-01-01 00:00:00') AS done_at,
								   t.is_deleted,
								   COALESCE(t.deleted_at, '0001-01-01 00:00:00') AS deleted_at
							FROM tasks t
									 JOIN task_priorities tp ON t.priority_id = tp.id
									 WHERE t.is_deleted = TRUE;`
)
