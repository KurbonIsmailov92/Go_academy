package main

import "fmt"

func main() {
	// phoneBook := map[string]string{
	// 	"Alice": "+992 93 774 77 30",
	// 	"Bob":   "+992 918 30 74 34",
	// }
	// fmt.Println(phoneBook, "\n")

	// for key, value := range phoneBook {
	// 	fmt.Println(key, value)
	// }

	// phoneBook["Vasya"] = "111 1111 111"

	// for key, value := range phoneBook {
	// 	fmt.Println(key, value)
	// }
	nums := []int{1,2,3,4,5,2}
	uniqueNums := make(map[int]int)

	for _, num := range nums {
		// _, ok := uniqueNums[num]
		// if ok == true {
		// 	uniqueNums[num]++

			
		// }else{
		// 	uniqueNums[num] = 1
		// }
		uniqueNums[num]++
	}





}
