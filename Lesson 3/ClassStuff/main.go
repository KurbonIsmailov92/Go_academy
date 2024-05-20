package main

import "fmt"

func main() {
	// DBlogin := "Vasya"
	// DBpassword := "vasya2024"
	// DBfullName := "Vasya Pupkin"
	// DBbalance := 100.00

	// fmt.Println("Input your login:")
	// var login string
	// fmt.Scan(&login)

	// fmt.Println("Input your password:")
	// var password string
	// fmt.Scan(&password)

	// if login == DBlogin && password == DBpassword {
	// 	fmt.Println("Hello", DBfullName)
	// 	fmt.Println("Your balance is", DBbalance)

	// 	fmt.Println("How much do you want to withdrawal?")
	// 	var withdrawalAmount float64
	// 	fmt.Scan(&withdrawalAmount)

	// 	if DBbalance < withdrawalAmount {
	// 		fmt.Println("You have not enough money")
	// 	}else {
	// 		DBbalance -= withdrawalAmount
	// 		fmt.Println("Your balance is", DBbalance)
	// 	}
	// }else {
	// 	fmt.Println("Wrong login or password")
	// }

	// var month int

	// fmt.Println("input name of the month: ")
	// fmt.Scan(&month)

	// switch month {
	// case 1, 2, 12:
	// 	fmt.Println("Winter")
	// case 3, 4, 5:
	// 	fmt.Println("Spring")
	// case 6, 7, 8:
	// 	fmt.Println("Summer")
	// case 9, 10, 11:
	// 	fmt.Println("Autumn")
	// default:
	// 	fmt.Println("There is no such season")

	// }

	var cathegory int

	fmt.Println("Choose the cathegory: 1 - Clothes, 2 - Shoes, 3 - Accesories")
	fmt.Scan(&cathegory)

	switch cathegory {
	case 1:
		fmt.Printf("Cost of Clothes is %v.", 15)
	case 2:
		fmt.Println("Cost of Shoes is 10")
	case 3:
		fmt.Println("Cost of Accesories is 5")
	default:
		fmt.Println("There is no such type of item!")

	}

}
