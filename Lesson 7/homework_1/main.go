package main

import (
	"fmt"
)

// Задача 1. Реверсирование map:
// Создайте программу, которая реверсирует ключи и значения в
// map. Предположим, что все значения уникальны.
// Input: {
//     "Vasya": "111-111",
//     "Alex": "222-222",
//     "Sadam": "333-333"
// }

// Output: {
//             "111-111": "Vasya",
//             "222-222": "Alex",
//             "333-333": "Sadam"
// }

func main() {
	phoneBook := make(map[string]string)
	reversedPhoneBook := make(map[string]string)
	var name, number string

	for {
		fmt.Println("Введите Имя:", "Для выхода из программы введите - exit:")
		fmt.Scanln(&name)

		if name != "exit" {
			_, exists := phoneBook[name]

			if exists {
				fmt.Printf("%s уже есть запись с таким именем \n", name)

				continue
			} else {
				fmt.Println("Введите номер: Например: (999-988-987)")
				fmt.Scanln(&number)

				phoneBook[name] = number

			}

		} else {
			for name, number := range phoneBook {
				_, exists := reversedPhoneBook[number]

				if exists {
					fmt.Printf("Номер %s для %s уже записан у %s \n", number, name, reversedPhoneBook[number])
					continue
				} else {
					reversedPhoneBook[number] = name
				}
			}
			fmt.Println("\nИзначальный список:")
			fmt.Println("===================================================")

			for name, number := range phoneBook {

				fmt.Println(name, ":", number)
			}
			fmt.Println("===================================================")
			fmt.Println("\nРезультат:")
			fmt.Println("===================================================")

			for number, name := range reversedPhoneBook {
				fmt.Println(number, ":", name)
			}
			fmt.Println("===================================================")		
		}
		break
	}
}
