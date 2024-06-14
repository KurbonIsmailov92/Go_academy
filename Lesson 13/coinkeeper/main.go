package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

type Operation struct {
	ID            int
	OperationType string
	Categorey     string
	Amount        float64
	CreatedAt     string
	Description   string
}

var Operations []Operation

func AddOperation(o Operation) {
	Operations = append(Operations, o)
}

func GetOperatoinsFromConsole() Operation {

	var o Operation
	o.ID = 1
	fmt.Println("Введите тип операции (income/outcome):")
	fmt.Scan(&o.OperationType)
	fmt.Scan(&o.Categorey)
	fmt.Scan(&o.Amount)
	//o.CreatedAt = "10-06-2024"
	//o.Description = "Начисление зарплаты"

	return o

}
