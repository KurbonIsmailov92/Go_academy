package main

import (
	"fmt"
	"strconv"
)

// Задача 1
// Подсчет количества цифр в числе:
// Запросите у пользователя число и определите,
// сколько цифр содержится в этом числе. Например,
// если пользователь ввел число 12345, программа должна
// сообщить, что в числе 12345 содержится 5 цифр.

func main() {
	var counter, forPrint int
	var input string

	fmt.Println("Введите число:")
	fmt.Scan(&input)

	// проверяем на то что ввод начинается с 0. Например: 01245 и исправляем на 1245
	if len(input) > 1 {

		runes := []rune(input)

		if string(runes[0]) == "0" {

			input = string(runes[1:])
			inputInt, _ := strconv.Atoi(input)

			forPrint = inputInt

			for inputInt > 0 {
				if inputInt%10 != 0 {
					counter++
				}
				inputInt /= 10
			}

		} else { // Если ввод не начинается с 0, то просто форматируем строку в цифры
			inputInter, _ := strconv.Atoi(input)
			forPrint = inputInter

			for inputInter > 0 {
				if inputInter%10 != 0 {
					counter++
				}
				inputInter /= 10
			}
		}

	} else { //Если длина ввода не больше одного, то работает этот блок
		inputInter, _ := strconv.Atoi(input)
		forPrint = inputInter
		counter++
	}

	//Для правильного окончания слова "цифра" испльзовал свитч
	switch counter {
	case 1:
		fmt.Printf("В числе %d содержится %d цифра", forPrint, counter)
	case 2, 3, 4:
		fmt.Printf("В числе %d содержится %d цифры", forPrint, counter)
	default:
		fmt.Printf("В числе %d содержится %d цифр", forPrint, counter)
	}
}
