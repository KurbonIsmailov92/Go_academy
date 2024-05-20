package main

import "fmt"

// Задача 1. Определение максимального числа
// Напишите программу, которая запрашивает у пользователя три целых числа и определяет, какое
// из них является наибольшим. Результат должен выводиться на экран.
// Пример:
// Введите три целых числа: 5 12 8
// Наибольшее число: 12

func main() {
	var first, second, third, maxInt int

	fmt.Println("Введите три целых числа:")
	fmt.Scan(&first, &second, &third)

	if first > second && first > third {
		maxInt = first
	} else if second > first && second > third {
		maxInt = second
	} else if third > first && third > second {
		maxInt = third
	}

	fmt.Printf("Наибольшее число: %d", maxInt)
}
