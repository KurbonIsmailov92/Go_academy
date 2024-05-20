package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Задача 4: Определение типа переменной
// Напишите программу, которая запрашивает у пользователя значение
// переменной (целое число, дробное число или строку) и определяет, какого
// типа это значение.
// Пример:
// Введите значение переменной: 3.14
// Тип переменной: дробное число

func main() {
	var input string
	flag := "string"

	fmt.Println("Введите значение переменной:")
	fmt.Scan(&input)

	if strings.Contains(input, ".") {
		if _, err := strconv.ParseFloat(input, 64); err == nil {
			flag = "float"
		}
	}

	if _, err := strconv.Atoi(input); err == nil {
		flag = "int"
	}

	switch flag {
	case "int":
		fmt.Println("Тип переменной: Целое число")
	case "float":
		fmt.Println("Тип переменной: Дробное число")
	default:
		fmt.Println("Тип переменной: Строка")
	}

}
