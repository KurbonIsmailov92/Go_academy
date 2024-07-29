package db

const (
	CreatingEmployeesTable = `CREATE TABLE employees
	(
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(30),
		last_name VARCHAR(30),
		age INT NOT NULL,
		department VARCHAR(30),
		position VARCHAR(30) NOT NULL DEFAULT 'Стажёр',
		salary INT,
		email VARCHAR(255) UNIQUE NOT NULL,
		is_deleted BOOLEAN NOT NULL DEFAULT FALSE
	);`
)
