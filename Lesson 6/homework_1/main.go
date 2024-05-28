package main

import (
	"fmt"
	"strings"
)

// Задача 1. Реверс слайса:
// Напишите программу, которая создает слайс строк и
// реверсирует его, то есть меняет порядок элементов на
// противоположный. Выведите на экран полученный
// реверсированный слайс.
// input: [1, 3, 4, 2, 6]
// output: [6, 2, 4, 3, 1]

func main() {
	var toReverse []string
	var input string

	fmt.Println("Программа заущена!")

	for {
		fmt.Println("Для продолжения работы введите символы:")
		fmt.Println("Для закрытия программы и вывода результата введите 'exit':")
		fmt.Scanln(&input)

		if strings.ToLower(input) != "exit" {
			toReverse = append(toReverse, input)
		} else {

			reversed := make([]string, len(toReverse))

			for i := 0; i < len(reversed); i++ {
				reversed[i] = toReverse[(len(toReverse)-1)-i]
			}
			fmt.Println("Реверс слайса", toReverse, "равна слайсу", reversed)
			fmt.Println("Программа закрыта!")
			break
		}

	}

}
