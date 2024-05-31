package main

import (
	"fmt"
)

func main() {

	//salaries := []int{1000, 2000, 3000, 500, 9000}
	// tax := 13

	// maxSalary:= MaxSalary(salaries)
	// fmt.Println(maxSalary)

	// taxed:= TaxCounter(maxSalary, tax)
	// fmt.Println(taxed)

	// phoneBook := map[string]string{
	// 	"111-111": "Vasya",
	// 	"222-222": "Bob",
	// 	"333-333": "Sadam",
	// 	"444-444": "Alice",
	// 	"555-555": "Vasya",
	// 	"666-666": "Andrey",
	// }

	//target:= "Vasya"

	// fmt.Println(findDublNames(target, phoneBook))



}

func findDublNames(taget string, phones map[string]string) (count uint)  {
	for _, v := range phones {
		if taget == v {
			count++
		}
	}
	return
}

func greet(name string) {
	fmt.Printf("Hello %s!", name)
}

func sum(a, b int) {
	fmt.Println(a + b)
}

// Find max element
func MaxSalary(arr []int) int {
	max := 0
	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}

func TaxCounter(num, tax int) int {
	num *= tax
	num /= 100
	return num
}
