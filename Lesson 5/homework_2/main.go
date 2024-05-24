package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Задача 2: Подсчет частоты элемента в массиве целых чисел
// Напишите программу, которая считывает с консоли целые числа и записывает их в массив. Затем
// подсчитайте частоту (количество повторений) заданного значения в массиве и выведите результат.
// // [0, 1, 2, 1, 3]
// input: 1
// output: Число 1 повторяется 2 раза

func main() {
	var numbers []int
	var input, forCompare, greeting, forLoop, forSearch string
	var count int

	greeting = "Программа запущена. Добро Пожаловать!"
	forLoop = "Для начала поиска повторений введите - search, для продолжения Введите число:"
	forSearch = "Введите число для поиска повторений:"

	fmt.Println(greeting)

	for {
		fmt.Println(forLoop)
		fmt.Scan(&input)
		input = strings.ToLower(input)

		if input == "search" {
			fmt.Println(forSearch)
			fmt.Scan(&forCompare)

			num, err := strconv.Atoi(forCompare)

			if err != nil {
				fmt.Println("Вы ввели букву!")
			} else {
				for _, value := range numbers {
					if value == num {
						count++
					}
				}

				fmt.Printf("Для числа %d найдено %d повторений", num, count)

				break

			}

		} else {
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Вы ввели букву!")
				continue
			} else {
				numbers = append(numbers, num)
			}

		}
	}

}
