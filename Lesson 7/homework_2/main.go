package main

import "fmt"

// Задача 2. Объединение двух map:
// Напишите программу, которая объединяет два map. В случае совпадения ключей,
// выберите значение из второго map.
// Input:
// map1: {
//     "Vasya": "111-111",
//     "Alex": "222-222",
//     "Sadam": "333-333",
// }
// map2: {
//     "Sadam": "444-444",
//     "Bob": "777-777",
// }

// Output: {
//             "Vasya": "111-111",
//             "Alex": "222-222",
//             "Sadam": "444-444",
//             "Bob": "777-777",
//         }

func main() {
	firstMap := make(map[string]string)
	secondMap := make(map[string]string)
	var name, number string

	for {
		fmt.Println("Введите Имя:", "Для ввода следующего списка введите - next:")
		fmt.Scanln(&name)

		if name != "next" {
			_, exists := firstMap[name]

			if exists {
				fmt.Printf("%s уже есть запись с таким именем \n", name)

				continue
			} else {
				fmt.Println("Введите номер: Например: (999-988-987)")
				fmt.Scanln(&number)

				firstMap[name] = number

			}

		} else {
			fmt.Sprintln("Первый список:")

			for name, number := range firstMap {
				fmt.Println(name, ":", number)
			}
			break
		}
	}

	result := firstMap

	for {
		fmt.Println("Введите Имя:", "Для выхода из программы введите - exit:")
		fmt.Scanln(&name)

		if name != "exit" {
			_, exists := secondMap[name]

			if exists {
				fmt.Printf("%s уже есть запись с таким именем \n", name)

				continue
			} else {
				fmt.Println("Введите номер: Например: (999-988-987)")
				fmt.Scanln(&number)

				secondMap[name] = number

			}

		} else {
			fmt.Println("\nПервый список:")

			for name, number := range firstMap {
				fmt.Println(name, ":", number)
			}

			fmt.Println("\nВторой список:")

			for name, number := range secondMap {
				fmt.Println(name, ":", number)
			}

			for name, number := range secondMap {
				_, exists := result[name]

				if exists {
					fmt.Printf("\nОбновлен номер %s для имени %s на новый %s!\n", result[name], name, number)
					result[name] = number
				} else {
					result[name] = number
				}
			}
			fmt.Println("\nРезультат:")

			for name, number := range result {
				fmt.Println(name, ":", number)
			}
			break
		}
	}

}
