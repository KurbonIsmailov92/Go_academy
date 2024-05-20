package main

import "fmt"

// Задача 2
// Вывод чисел в обратном порядке:
// Запросите у пользователя число и выведите все
// числа от этого числа до 1 в обратном порядке.
// 5
// 5 4 3 2 1

func main() {
	var input int

	fmt.Println("Введите число:")
	fmt.Scan(&input)

	if input > 0 {
		for input > 0 {
			fmt.Print(input, " ")
			input--
		}
	} else {
		for input <= 0 {
			fmt.Print(input, " ")
			input++
		}
	}
}
