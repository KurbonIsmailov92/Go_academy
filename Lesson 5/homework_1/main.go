package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Задача 1: Найти среднее значение элементов массива дробных чисел
// Напишите программу, которая запрашивает у пользователя ввод дробных чисел (float64)
// для заполнения массива и затем вычисляет и выводит среднее значение элементов массива.
// summa / count

func main() {
	var numbers []float64
	var sum, avg float64
	var input, greeting, forLoop, message, ending string

	greeting = "Программа запущена. Добро Пожаловать!"
	forLoop = "Для выхода из программы, введите - exit, для продолжения Введите число:"
	ending = "Вы закрыли программу"

	fmt.Println(greeting)

	for {

		fmt.Println(forLoop)
		fmt.Scan(&input)
		input = strings.ToLower(input)

		if input == "exit" {
			fmt.Println(message)
			fmt.Println(ending)

			break
		} else {
			num, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Вы ввели букву!")
				continue
			} else {
				numbers = append(numbers, num)
				sum += num
				avg = sum / float64(len(numbers))
				message = fmt.Sprintf("Сумма: %.2f Количество: %d Средняя: %.2f", sum, len(numbers), avg)
				fmt.Println(message)
			}
		}

	}
}




