package main

import "fmt"

func main() {
	sumOfPrimes()

}

func IsPrime(number int) int {

	if number <= 1 {
		return 0
	}

	for i := 2; i*i <= number; i++ { // 81
		if number%i == 0 {
			return 0
		}
	}
	return number
}

func sumOfPrimes(){
	var first, second, third, sum int

	fmt.Scan(&first, &second, &third)

	first = IsPrime(first)
	second = IsPrime(second)
	third = IsPrime(third)

	sum = first + second + third

	fmt.Println(sum)

	if IsPrime(sum) != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}