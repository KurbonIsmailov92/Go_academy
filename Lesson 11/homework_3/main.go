package main

import "fmt"

// Задача 3: Указатели на слайсы
// Напишите функцию, которая принимает указатель на слайс целых чисел и удваивает
// каждый элемент в слайсе.

func main() {
	numbers := []int{2, 5, 4, 6, 9}
	fmt.Println("Числа до:", numbers)

	dubleOfInts(numbers)
	fmt.Println("Числа после:", numbers)

}

func dubleOfInts(numbers []int) {

	for i := 0; i < len(numbers); i++ {
		numbers[i] *= 2
	}

}
