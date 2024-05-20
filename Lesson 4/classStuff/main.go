package main

import (
	"fmt"
)

func main() {

	// for i := 1; i <= 10; i += 2 {
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 10; i++ {

	// 	if i%2 == 0 {
	// 		continue
	// 	}else {
	// 		fmt.Println(i)
	// 	}

	// }

	// for i := 1; i <= 100; i++ {
	// 	if i% 5 ==0 {
	// 		fmt.Println(i, "Yes")
	// 		break
	// 	}else{
	// 		fmt.Println(i, "No")
	// 	}
	// }

	// i := 0
	// for i<=10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// var input int
	// fmt.Scan(&input)
	// sum:=0
	// for i := 0; i <= input; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum).

	// var number int
	// fmt.Println("Input number")
	// fmt.Scan(&number)

	// for i := 1; i <= 10; i++ {
	// 	sum:= 0
	// 	sum = number * i
	// 	fmt.Println(number, "x", i, "=", sum )
	// }

	// var num, count int
	// fmt.Scan(&num)

	// for i := 1; i <= num; i++ {
	// 	if num%i == 0 {
	// 		count++
	// 	}

	// }

	// if count == 2 {
	// 	fmt.Println("Simple")
	// } else {
	// 	fmt.Println("Not Siple")
	// }

	var num, sum int
	fmt.Println("Input num:")
	fmt.Scan(&num)

	for num >= 0 {
		sum += num
		fmt.Println("Sum is", sum)
		fmt.Println("Input num:")
		fmt.Scan(&num)
	}
	fmt.Println("Sum is:", sum)

}
