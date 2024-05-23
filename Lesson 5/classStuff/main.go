package main

import (
	"fmt"
	"strconv"
	//"golang.org/x/text/number"
)

// s1 := 100
// s2 := 100
// s3 := 100
// s4 := 100
// s5 := 100

func main() {
	// var numbers [100] int
	// var numes = [5]int{1500, 2000, 3200, 4400, 5400}

	// for i, _ := range numbers {

	// 	numbers[i] = i + 1

	// }

	// for index, value := range numes {
	// 	fmt.Println( index, value)
	// }

	// // var numes2 = [...]int{1, 2, 3, 4, 5}

	// fmt.Println(len(numes), cap(numes))

	// var numes2 = [5]int{1, 2}

	// fmt.Println(len(numes2), cap(numes2))

	// fmt.Println(numbers)

	// var salaries [...]float32

	// for i := range salaries {
	// 	fmt.Println("Input slary of worker")
	// 	fmt.Scan(&salaries[i])
	// }

	// // for i, v := range salaries {
	// // 	fmt.Println("Worker", i+1, "salary", v)
	// // }

	// var sum float32
	// for _, value := range salaries {
	// 	sum += value
	// }

	// fmt.Println("Sum is", sum)

	var numbers []float64
	var num, sum, avg float64
	var input, message string

	fmt.Println("Программа запущена. Добро Пожаловать!")

	for {

		fmt.Println("Для выхода из программы, введите - exit, для продолжения Введите число:")
		fmt.Scan(&input)

		if input == "exit" {
			fmt.Println(message)
			fmt.Println("Вы закрыли программу")

			break
		} else {
			num, _ = strconv.ParseFloat(input, 64)
			numbers = append(numbers, num)
			sum += num
			avg = sum / float64(len(numbers))
			message = fmt.Sprintf("Сумма: %.2f Количество: %d Средняя: %.2f", sum, len(numbers), avg)
			fmt.Println(message)
		}

	}

}
