package main

import "fmt"

func main() {
	// first := []int{1, 2, 3, 4}
	// second := []int{2, 5, 1}

	// uniqMap := make(map[int]int)
	// temp := append(first, second...)

	// for _, value := range temp {
	// 	uniqMap[value]++

	// }

	// for key, value := range uniqMap {
	// 	if value == 1 {
	// 		fmt.Print(key, ";")
	// 	}
	// }
	//-----------------------------------------------------//
	// first := map[string]int{
	// 	"Alex":  3000,
	// 	"Bob":   2000,
	// 	"Vasya": 1000,
	// 	"Sadam": 500}

	// second := map[string]int{
	// 	"Anna":  3000,
	// 	"Vasya": 1000,
	// 	"Sadam": 2000,
	// }

	// result := make(map[string]int)

	// for key, value := range first {
	// 	value2, ok := second[key]

	// 	if ok {
	// 		result[key] = value + value2
	// 	}
	// }
	// for key, value := range result {
	// 	fmt.Println(key, ":", value)
	// }
	//----------------------------------------------------//

	fruits := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"cherry": "yellow",
		"lemon":  "yellow",
		"grape":  "purple",
	}

	for key, value := range fruits {
		target := "yellow"
		if value == target {
			fmt.Println(key)
		}
	}

}
