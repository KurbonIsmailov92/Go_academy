package main

import "fmt"

// s1 := 100
// s2 := 100
// s3 := 100
// s4 := 100
// s5 := 100

func main() {
	// var nums [100] int
	// var numes = [5]int{1500, 2000, 3200, 4400, 5400}

	// for i, _ := range nums {

	// 	nums[i] = i + 1

	// }

	// for index, value := range numes {
	// 	fmt.Println( index, value)
	// }

	// // var numes2 = [...]int{1, 2, 3, 4, 5}

	// fmt.Println(len(numes), cap(numes))

	// var numes2 = [5]int{1, 2}

	// fmt.Println(len(numes2), cap(numes2))

	// fmt.Println(nums)

	var salaries [10]float32

	for i := range salaries {
		fmt.Println("Input slary of worker")
		fmt.Scan(&salaries[i])
	}

	// for i, v := range salaries {
	// 	fmt.Println("Worker", i+1, "salary", v)
	// }

	var sum float32
	for _, value := range salaries {
		sum += value
	}

	fmt.Println("Sum is", sum)

}
