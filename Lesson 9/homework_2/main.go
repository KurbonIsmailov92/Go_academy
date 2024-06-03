package main

import "fmt"

// Задача 2: Функция для фильтрации слайса
// Напишите функцию, которая принимает слайс чисел и возвращает новый слайс, содержащий
// только четные числа.
// input: [2, 5, 8, 10, 9]
// output: [2, 8, 10]

func main() {
	exampleSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(onlyEvenNumbers(exampleSlice))

}

func onlyEvenNumbers(slice []int) []int {
	var forReturn []int

	for _, value := range slice {

		if value%2 == 0 {
			forReturn = append(forReturn, value)
		}
	}

	return forReturn
}
