package main

import (
	"crud_from_go/db"
	"crud_from_go/models"
	"crud_from_go/repository"
	"fmt"
)

func main() {
	err := db.ConnectToDB()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	defer db.CloseDBConnection()

	petya := models.Employee{
		FirstName:  "Petya",
		LastName:   "Petrov",
		Age:        24,
		Department: "Бухгалтерия",
		Position:   "Бухгалтер",
		Salary:     12650,
		Email:      "petya@gmail.com",
		IsDeleted:  false,
	}
	masha := models.Employee{
		FirstName:  "Masha",
		LastName:   "Ivanova",
		Age:        26,
		Department: "Приёмная",
		Position:   "Секретрь",
		Salary:     24000,
		Email:      "nashamasha@gmail.com",
		IsDeleted:  false,
	}

	err = repository.CreateEmployee(petya)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	err = repository.CreateEmployee(masha)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	employees, err := repository.GetAllEmployees(false)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	fmt.Println(employees)

	employee, err := repository.GetEmployeeID(1)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	fmt.Println(employee)

	employeesForUpdate := models.Employee{
		ID:         17,
		Department: "Хозотдел",
		Salary:     1200,
	}

	err = repository.UpdateEmployeeByID(employeesForUpdate)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	err = repository.SoftDeleteEmployeeByID(17)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	err = repository.HardDeleteEmployeeByID(16)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

}
