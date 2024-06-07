package main

import "fmt"

// Задача 4: Возвращение указателя из функции:
// Напишите функцию, которая принимает два целых числа и возвращает указатель на большее из них.

// Задача 5: Разыменование указателя:
// Напишите функцию, которая принимает два указателя на целые числа и возвращает их разность.

// Задача 6: Указатели и карты:
// Напишите функцию, которая принимает карту "ключ-значение" (где значения - целые числа) и увеличивает значение
// каждого ключа на 1, используя указатели.

func main() {

	// Задача 4
	fmt.Println(returnPointer(3, 8))

	// Задача 5
	var first, second int
	first = 8
	second = 3
	fmt.Println(subtract(&first, &second))

	//// Задача 6
	salaries := map[string]int{
		"Kolya": 2400,
		"Vasya": 4400,
		"Petya": 1300,
	}

	mapDuble(salaries)
	fmt.Println(salaries)
}

// Задача 4: Возвращение указателя из функции
func returnPointer(first, second int) (ptrMax *int) {
	if first > second {
		return &first
	}
	if first < second {
		return &second
	} else {
		fmt.Println("Оба числа равны")
		return
	}
}

// Задача 5: Разыменование указателя:
func subtract(ptrFirst, ptrSecond *int) int {
	return *ptrFirst - *ptrSecond
}

// Задача 6: Указатели и карты:
func mapDuble(randomMap map[string]int) {
	for key, value := range randomMap {
		randomMap[key] = value * 2
	}
}
