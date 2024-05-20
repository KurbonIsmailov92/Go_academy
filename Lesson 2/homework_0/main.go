package main

import "fmt"

// Задача 0: Расчет стоимости билета
// Напишите программу, которая запрашивает у пользователя его возраст и выводит стоимость
// билета в кинотеатр в соответствии с возрастной категорией:
// Дети до 7 лет: бесплатно
// Дети от 7 до 18 лет: 50% скидка
// Взрослые: полная стоимость
// Пенсионеры: 25% скидка

func main() {
	var usersAge uint16
	ticketsCost := 100
	fmt.Println("Салом, введите свой возраст:")
	fmt.Scan(&usersAge)

	if usersAge < 7 {
		fmt.Println("У вас бесплатный билет!")
	} else if usersAge >= 7 && usersAge < 18 {
		ticketsCost -= ticketsCost / 2
		fmt.Printf("Стоимость билета, с учетом скидки в 50 процентов, для Вас будет составлять %d сомони", ticketsCost)
	} else if usersAge > 65 {
		ticketsCost -= ticketsCost / 4
		fmt.Printf("Стоимость билета, с учетом скидки в 25 процентов, для Вас будет составлять %d сомони", ticketsCost)
	} else {
		fmt.Printf("Стоимость билета для Вас будет составлять %d сомони", ticketsCost)
	}

}
