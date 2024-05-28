package main

import "fmt"

// Задача 2. Объединение слайсов без дубликатов:
// Напишите программу, которая создает два слайса строк
// и объединяет их без дубликатов, то есть исключая
// повторяющиеся элементы. Выведите на экран результат
// объединения.
// input: [1, 2, 3], [3, 4, 5]
// output: [1, 2, 4, 5]

func main() {
	var firstSlice, secondSlice, result []string
	var input string

	fmt.Println("Для заполнения первого слайса, введите данные:")

	for {

		fmt.Scan(&input)

		if input != "next" {

			firstSlice = append(firstSlice, input)

			fmt.Println("Для перехода к заполнению второго слайса введите - next \n", "Для продолжение введите данные:")
		} else {
			fmt.Println("Первый слайс:", firstSlice)
			break
		}

	}

	fmt.Println("Для заполнения второго слайса, введите данные:")
	result = firstSlice

	for {

		fmt.Scan(&input)

		if input != "result" {
			secondSlice = append(secondSlice, input)
			fmt.Println("Для вывода результата введите - result \n", "Для продолжение введите данные:")
		} else {
			fmt.Println("Первый слайс:", firstSlice, "Второй слайс:", secondSlice)

			for i := 0; i < len(secondSlice); i++ {
				flag := true
				for j := 0; j < len(result); j++ {
					if secondSlice[i] == result[j] {
						flag = false
						result = append(result[:j], result[j+1:]...)
					}
				}

				if flag {
					result = append(result, secondSlice[i])
				}

			}

			fmt.Println("Результат:", result)
			fmt.Println("Программа закрыта!")
			break
		}
	}
}
