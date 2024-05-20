package main

import "fmt"

// Задача 1: Калькулятор
// Напишите программу калькулятор, которая запрашивает у пользователя два
// числа и оператор (сложение, вычитание, умножение, деление) и выводит
// результат операции.
// Пример:
// Введите первое число: 10
// Введите второе число: 5
// Выберите оператор (+, -, *, /): *
// Результат: 50

func main() {
	var firstNumber, secondNumber, result int
	var operator string

	fmt.Println("Введите первое число:")
	fmt.Scan(&firstNumber)

	fmt.Println("Введите второе число:")
	fmt.Scan(&secondNumber)

	fmt.Println("Выберите оператор (+, -, *, /):")
	fmt.Scan(&operator)

	if operator == "/" && secondNumber == 0 {
		fmt.Println("Нельзя делить на ноль!")
	} else {
		switch operator {
		case "+":
			result = firstNumber + secondNumber
			fmt.Println("Результат:", result)
		case "-":
			result = firstNumber - secondNumber
			fmt.Println("Результат:", result)
		case "*":
			result = firstNumber * secondNumber
			fmt.Println("Результат:", result)
		case "/":
			result = firstNumber / secondNumber
			fmt.Println("Результат:", result)
		default:
			fmt.Println("Нет такого действия!")
		}
	}

}
