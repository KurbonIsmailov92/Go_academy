package main

import "fmt"

// Задача 3: Объединение значений по ключам
// Напишите функцию, которая принимает две карты с одинаковыми ключами и объединяет их значения в новую карту.
// Если ключ присутствует в обеих картах, сложите их значения.
// Input:
// map1: {
//     "Vasya": 123,
//     "Alex": 3124,
//     "Sadam": 124,
// }
// map2: {
//     "Sadam": 124,
//     "Bob": 145,
// }

// Output: {
//             "Vasya": 123,
//             "Alex": 3124,
//             "Sadam": 248,
//             "Bob": 145,
//         }

func main() {
	firstList := map[string]int{
		"Vasya": 1,
		"Alex":  1,
		"Sadam": 1}

	secondList := map[string]int{
		"Sadam": 1,
		"Bob":   1,
	}

	fmt.Println(mergeMaps(firstList, secondList))

}

func mergeMaps(firsMap, secondMap map[string]int) (resultMap map[string]int) {

	resultMap = firsMap

	for key, value := range secondMap {
		_, ok := resultMap[key]

		if ok {
			resultMap[key] += value
		} else {
			resultMap[key] = value
		}

	}

	return
}
