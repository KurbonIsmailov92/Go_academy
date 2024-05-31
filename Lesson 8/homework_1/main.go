package main

import "fmt"

// Задача 1: Поиск отсутствующих ключей
// Напишите программу, которая принимает карту и список ключей. Программа должна вернуть список ключей,
// которые отсутствуют в карте.
// input:
// {
//         "apple":  1,
//         "banana": 2,
//         "cherry": 3,
// }

// []string{"banana", "grape", "apple", "pineapple"}

// output: grape, pineapple

func main() {

	autoList := map[string]int{
		"mersedes": 1,
		"audi":     2,
		"bmw":      3,
		"kia":      4,
		"toyota":   5,
	}

	searchBrends := []string{"honda", "bmw", "kia", "opel", "audi", "ford", "toyota"}

	for _, brend := range searchBrends {
		_, ok := autoList[brend]
		if !ok {
			fmt.Print(brend, ", ")
		}
		continue
	}
	//"honda", "opel", "ford"
}
