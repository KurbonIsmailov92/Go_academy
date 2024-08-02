package db

const (
	CreateUsersTableQuery = `CREATE TABLE IF NOT EXISTS users
							(
								id   SERIAL PRIMARY KEY,
								name VARCHAR(30)
							);`
	CreateTaskPrioritiesTableQuery = `CREATE TABLE IF NOT EXISTS task_priorities
								(
									id       SERIAL PRIMARY KEY,
									priority VARCHAR(20)
								);`
	CreateTasksTableQuery = `CREATE TABLE IF NOT EXISTS tasks
							(
								id               SERIAL PRIMARY KEY,
								title            VARCHAR NOT NULL,
								description      TEXT,
								task_priority_id INT REFERENCES task_priorities (id) NOT NULL,
								user_id          INT REFERENCES users (id) NOT NULL ON DELETE CASCADE,
								created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
								is_done          BOOLEAN   DEFAULT FALSE,
								done_at          TIMESTAMP DEFAULT NULL,
								is_deleted       BOOLEAN   DEFAULT FALSE,
								deleted_at       TIMESTAMP DEFAULT NULL
							)`
)
