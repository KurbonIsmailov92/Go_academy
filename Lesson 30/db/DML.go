package db

const (
	CreateEmployeeQuery         = "INSERT INTO employees (first_name, last_name, age, department, position, salary, email) VALUES ($1, $2, $3, $4, $5, $6, $7);"
	GetAllActiveEmployeesQuery  = "SELECT id, first_name, last_name, age, department, position, salary, email, is_deleted FROM employees WHERE is_deleted = $1;"
	GetEmployeeByIDQuery        = "SELECT id, first_name, last_name, age, department, position, salary, email, is_deleted FROM employees WHERE id = $1;"
	UpdateEmployeeInfoByIDQuery = "UPDATE employees SET department = $1, salary = $2 WHERE id = $3;"
	SoftDeleteEmployeeByIDQuery = "UPDATE employees SET is_deleted = TRUE WHERE id = $1;"
	HardDeleteEmployeeByIDQuery = "DELETE FROM employees WHERE id = $1;"
)
