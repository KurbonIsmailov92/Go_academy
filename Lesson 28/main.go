package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := ConnectToDB()
	if err != nil {
		panic(err)
	}

	defer CloseDBConnection(db)

	allEmployees, err := GetAllEmployees(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("All Employees:", allEmployees)

	chosenEmloyeeByID, err := GetEmployeeByID(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Chosen Employee is:", chosenEmloyeeByID)
}

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
	Age        int
	Position   string
}

func ConnectToDB() (*sql.DB, error) {
	connectionString := "user=kurbon password=ismoil dbname=online_store_db sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to database")

	return db, nil
}

func CloseDBConnection(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	fmt.Println("Closed database connection")
	return nil
}

func GetAllEmployees(db *sql.DB) ([]Employee, error) {
	rows, err := db.Query("SELECT id, name, department, salary, age, position FROM employees;")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var employees []Employee

	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.ID, &e.Name, &e.Department, &e.Salary, &e.Age, &e.Position)
		if err != nil {
			fmt.Println(err)
			continue
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func GetEmployeeByID(db *sql.DB) (Employee, error) {
	var id int
	fmt.Println("Enter employees ID:")
	fmt.Scan(&id)
	var e Employee
	row := db.QueryRow("SELECT id, name, department, salary, age, position FROM employees WHERE id = $1", id)
	err := row.Scan(&e.ID, &e.Name, &e.Department, &e.Salary, &e.Age, &e.Position)
	if err != nil {
		return Employee{}, err
	}

	return e, nil
}
