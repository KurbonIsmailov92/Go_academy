package main

import "fmt"

// Задача 3*
// Проверка числа на палиндром:
// Запросите у пользователя число и определите, является ли оно
// палиндромом. Палиндромом считается число, которое читается
// одинаково как справа налево, так и слева направо.
// Например, 121, 12321.

func main() {
	var input, forCompare, modifiedInput uint

	fmt.Println("Введите число:")
	fmt.Scan(&input)

	insaneInput := input
	forCompare = input

	for forCompare > 0 {
		forCompare %= 10
		modifiedInput += forCompare
		input /= 10
		forCompare = input
		modifiedInput *= 10
	}

	modifiedInput /= 10

	if insaneInput == modifiedInput {
		fmt.Println("Введенное число является палиндромом")
	} else {
		fmt.Println("Введенное число НЕ является палиндромом")
	}

}
