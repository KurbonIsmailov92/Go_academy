package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numbers []float64
	var num, sum, avg float64
	var input, message string

	fmt.Println("Программа запущена. Добро Пожаловать!")

	for {

		fmt.Println("Для выхода из программы, введите - exit, для продолжения Введите число:")
		fmt.Scan(&input)

		if input == "exit" {
			fmt.Println(message)
			fmt.Println("Вы закрыли программу")

			break
		} else {
			num, _ = strconv.ParseFloat(input, 64)
			numbers = append(numbers, num)
			sum += num
			avg = sum / float64(len(numbers))
			message = fmt.Sprintf("Сумма: %.2f Количество: %d Средняя: %.2f", sum, len(numbers), avg)
			fmt.Println(message)
		}

	}
}
