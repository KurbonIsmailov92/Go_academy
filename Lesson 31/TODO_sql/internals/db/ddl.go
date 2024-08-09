package db

const (
	CreateTaskPrioritiesTableQuery = `CREATE TABLE IF NOT EXISTS task_priorities
								(
									id       SERIAL PRIMARY KEY,
									priority VARCHAR(20)
								);`
	CreateTasksTableQuery = `CREATE TABLE IF NOT EXISTS tasks
							(
								id               SERIAL PRIMARY KEY,
								title            VARCHAR NOT NULL,
								description      VARCHAR(30),
								task_priority_id INT REFERENCES task_priorities (id) NOT NULL,
								created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
								user_name        VARCHAR(12) NOT NULL,
								is_done          BOOLEAN   DEFAULT FALSE,
								done_at          TIMESTAMP DEFAULT NULL,
								is_deleted       BOOLEAN   DEFAULT FALSE,
								deleted_at       TIMESTAMP DEFAULT NULL
							)`
)
