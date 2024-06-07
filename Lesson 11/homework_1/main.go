package main

import "fmt"

// Задача 1: Увеличение значения по указателю
// Напишите функцию, которая принимает указатель на целое число и увеличивает
// значение этого числа на заданное количество.

func main() {
	var intNumber, multiplyer int

	fmt.Println("Введите целое число:")
	fmt.Scan(&intNumber)

	fmt.Println("Введите множитель:")
	fmt.Scan(&multiplyer)

	fmt.Println("Результат:", multiply(&intNumber, multiplyer))

}

func multiply(ptrIntNumber *int, multiplyer int) int {
	*ptrIntNumber *= multiplyer
	return *ptrIntNumber
}
