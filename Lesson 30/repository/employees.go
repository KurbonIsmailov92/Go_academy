package repository

import (
	"crud_from_go/db"
	"crud_from_go/models"
	"fmt"
)

func CreateEmployee(e models.Employee) error {
	_, err := db.GetDBConnection().Exec(db.CreateEmployeeQuery,
		e.FirstName, e.LastName, e.Age, e.Department, e.Position, e.Salary, e.Email)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return err
	}

	return nil
}

func GetAllEmployees(isAvailable bool) ([]models.Employee, error) {
	rows, err := db.GetDBConnection().Query(db.GetAllActiveEmployeesQuery, isAvailable)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var e models.Employee
		err = rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Age, &e.Department, &e.Position, &e.Salary, &e.Email, &e.IsDeleted)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func GetEmployeeID(id int) (models.Employee, error) {
	var e models.Employee
	row := db.GetDBConnection().QueryRow(db.GetEmployeeByIDQuery, id)
	err := row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Age, &e.Department, &e.Position, &e.Salary, &e.Email, &e.IsDeleted)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return models.Employee{}, err
	}

	return e, nil
}

func UpdateEmployeeByID(e models.Employee) error {
	_, err := db.GetDBConnection().Exec(db.UpdateEmployeeInfoByIDQuery, e.Department, e.Salary, e.ID)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return err
	}
	return nil
}

func HardDeleteEmployeeByID(id int) error {
	_, err := db.GetDBConnection().Exec(db.HardDeleteEmployeeByIDQuery, id)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return err
	}

	return nil
}

func SoftDeleteEmployeeByID(id int) error {
	_, err := db.GetDBConnection().Exec(db.SoftDeleteEmployeeByIDQuery, id)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return err
	}
	return nil
}
