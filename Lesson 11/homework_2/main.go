package main

import "fmt"

// Задача 2: Поменять местами два значения
// Напишите функцию, которая принимает два указателя на целые числа и меняет их
// значения местами.

func main() {
	var first, second int

	fmt.Println("Введите первое чило:")
	fmt.Scan(&first)

	fmt.Println("Введите второе чило:")
	fmt.Scan(&second)

	fmt.Println("Первое значение:", first, "Второе значение:", second)

	replaceOfIntegers(&first, &second)

	fmt.Println("Результат: Первое значение:", first, "Второе значение:", second)

}

func replaceOfIntegers(ptrFirstInt, ptrSecondInt *int) {
	*ptrFirstInt, *ptrSecondInt = *ptrSecondInt, *ptrFirstInt
}
