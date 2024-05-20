package main

import "fmt"

// 1. Инициализация переменных:
// Задача: Объявите и проинициализируйте три переменные типа int - x, y и z - значениями,
// введенными с консоли. Выведите их на экран.

func main() {
	var x, y, z int

	fmt.Println("Введите значение для X:")
	fmt.Scan(&x)

	fmt.Println("Введите значение для Y:")
	fmt.Scan(&y)

	fmt.Println("Введите значение для Z:")
	fmt.Scan(&z)

	fmt.Println("Значение X:", x, "Значение Y:", y, "Значение Z:", z)
}
