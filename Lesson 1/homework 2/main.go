package main

import "fmt"

// 2. Использование нескольких переменных в одном выражении:
// Задача: Введите значения переменных a, b и c с консоли.
// Посчитайте сумму a + b + c и выведите ее на экран.

func main() {
	var a, b, c, result int

	fmt.Println("Введите значения А, Б, С: ")
	fmt.Scan(&a, &b, &c)

	result = a + b + c

	fmt.Println("Сумма введенных чисел равна:", result)
}
