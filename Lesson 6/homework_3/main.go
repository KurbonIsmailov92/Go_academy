package main

import (
	"fmt"
	"strconv"
)

// Задача 3. Циклический сдвиг слайса:
// Напишите программу, которая создает слайс целых чисел и
// выполняет циклический сдвиг его элементов вправо на
// заданное количество позиций. Например, если слайс
// содержит элементы [1, 2, 3, 4, 5] и выполнить циклический
// сдвиг на 2 позиции вправо, то результатом будет слайс
// [4, 5, 1, 2, 3]. Выведите на экран полученный результат.
// input: [1, 2, 3, 4, 5]
// output: [5, 1, 2, 3, 4]

func main() {

	var numbers []int
	var input string
	var index int

	fmt.Println("Салом. Введите число для добавления в слайс")

	for {
		fmt.Scan(&input)
		if input != "result" {
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Вы ввели букву!")
			} else {
				numbers = append(numbers, num)
			}
			fmt.Println("Для введения Коэффициента сдвига и показа результата введите - result")
			fmt.Println("Для продолжения введите число:")
		} else {
			fmt.Println("Слайс:", numbers)
			fmt.Printf("Введите Коэффициент сдвига элементов слайса от 0 до %d:", len(numbers)-1)
			fmt.Scan(&index)

			sliceLen := len(numbers)

			if index > sliceLen {
				fmt.Println("Ошибка! Введено число вне указанного диапазона")
				break
			}

			if index == 0 {
				fmt.Println("Результат:", numbers)
				break
			} else {
				numbers = append(numbers[sliceLen-index:], numbers[:sliceLen-index]...)

				fmt.Println(index)

				fmt.Println("Результат:", numbers)
				break
			}
		}

	}

}
