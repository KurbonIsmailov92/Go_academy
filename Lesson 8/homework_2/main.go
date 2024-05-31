package main

import "fmt"

// Задача 2: Удаление ключей по значению
// Напишите программу, которая принимает карту и значение. Программа должна удалить все ключи,
// соответствующие данному значению, и вернуть обновленную карту.
// input:
// {
//         "apple":  1,
//         "banana": 2,
//         "cherry": 2,
//         "date":   1,
// }

// valueToRemove := 1

// output:
// {
//         "banana": 2,
//         "cherry": 2,
// }

func main() {
	studentList := map[string]string{
		"Harry":   "Griffindor",
		"Drago":   "Slytherin",
		"Ron":     "Griffindor",
		"Cedric":  "Hufflepuff",
		"Hemione": "Griffindor",
		"Luna":    "Ravenclaw",
	}

	valueToRemove := "Griffindor"

	for name, hogwartsHouse := range studentList {
		if hogwartsHouse == valueToRemove {
			delete(studentList, name)
		}
	}

	fmt.Println("--------------------------------")

	for name, hogwartsHouse := range studentList {
		fmt.Println(name, hogwartsHouse)
	}

	fmt.Println("--------------------------------")

}
