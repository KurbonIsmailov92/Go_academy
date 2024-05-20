package main

import "fmt"

func main() {

	/*	age := 16
		var isAdult bool
		isAdult = age >= 18
		fmt.Println(isAdult)
	*/
	// fmt.Println("Enter your age:")
	// var age uint
	// fmt.Scan(&age)

	// if age >= 18 {
	// 	fmt.Println("You are adult")
	// }
	// if age < 18 {
	// 	fmt.Println("You are teen")
	// }
	// if age > 7 {
	// 	fmt.Println("You are todler")
	// }
	// 	fmt.Println("You are baby")

	// var number int
	// fmt.Println("Enter the number:")
	// fmt.Scan(&number)
	// if number%2 == 0 {
	// 	fmt.Println("Четное")
	// } else {
	// 	fmt.Println("Нечетное")
	// }

	// var hour uint
	// fmt.Println("Enter hour")
	// fmt.Scan(&hour)

	// if hour > 24 {
	// 	fmt.Println("Net takogo chasa")
	// }else if hour >= 4 && hour < 12 {
	// 	fmt.Println("Morning")
	// }else if hour == 12 {
	// 	fmt.Println("Noon")
	// }else if hour > 12 && hour < 20 {
	// 	fmt.Println("Evening")
	// }else if hour == 0 {
	// 	fmt.Println("Midnight")
	// }else{
	// 	fmt.Println("Night")
	// }

	// var year uint
	// fmt.Scan(&year)
	// if year%4 == 0 && year%100 != 0 || year%400 == 0 {
	//     	fmt.Println("vis")
	// } else {
	// 	fmt.Println("Nevis")
	// }

	var age uint16
	fmt.Println("Enter age:")
	fmt.Scan(&age)

	if age > 0 && age <= 6 {
		fmt.Println("Doshkolnik")
	} else if age >= 7 && age <= 17 {
		fmt.Println("Shkolnik")
	} else if age >= 18 && age <= 24 {
		fmt.Println("Student")
	} else {
		fmt.Println("Adult")
	}

}
