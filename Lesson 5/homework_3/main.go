package main

// Задача 3: Поиск минимального и максимального элементов в массиве вещественных чисел
// Напишите программу, которая считывает с консоли вещественные числа и записывает их в
// массив. Затем найдите и выведите минимальный и максимальный элементы а также среднее
// значение элементов в этом массиве.
// // [0, 1, 2, 3]
// // output:
// min: 0
// max: 3
// avg: 1.5

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numbers []float64
	var sum, avg, min, max float64
	var input, greeting, forLoop, message, ending string

	greeting = "Программа запущена. Добро Пожаловать!"
	forLoop = "Для выхода из программы, введите - result, для продолжения Введите число:"
	ending = "Вы закрыли программу"

	fmt.Println(greeting)

	for {

		fmt.Println(forLoop)
		fmt.Scan(&input)
		input = strings.ToLower(input)

		if input == "result" {
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

				for _, value := range numbers {
					if value > max {
						max = value
					} else {
						continue
					}
				}

				min = max

				for _, value := range numbers {
					if value < min {
						min = value
					} else {
						continue
					}

				}

				message = fmt.Sprintf("Минимум: %.2f Максимум: %2.f Средняя: %.2f", min, max, avg)

			}
		}

	}
}
