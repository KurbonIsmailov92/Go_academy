package main

import "fmt"

// Задача 1: Напиши функцию, которая принимает слайс чисел и возвращает их
// произведение.
// input: [2, 5, 6]
// output: 60

func main() {
	example := []int{1, 2, 3, -2, 4, -1}

	fmt.Println(sumOfSliceElements(example)) // 48

}

//Функция sumOfSliceElements принимает слайс из чисел и возвращает результат произведения всех его элементов
func sumOfSliceElements(slice []int) int {
	sum := 1
	for _, value := range slice {
		sum *= value
	}

	return sum
}
