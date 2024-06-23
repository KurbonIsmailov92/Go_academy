package main

import "fmt"

/*
Задача 3: Интерфейс с множественными методами
Создайте интерфейс Person с методами GetName() string и GetAge()
int. Затем создайте структуру Student, которая реализует методы
интерфейса Person. Напишите функцию printPersonInfo, которая
принимает интерфейс Person и выводит имя и возраст.
*/

func main() {
	Eric := Student{
		Name: "Eric Cartman",
		Age:  18,
	}

	Kyle := Student{
		Name: "Kyle Broflovski",
		Age:  20,
	}

	Stan := Student{
		Name: "Stan Marsh",
		Age:  19,
	}

	Kenny := Student{
		Name: "Kenny McCormick",
		Age:  21,
	}

	printPerson(Eric)
	printPerson(Kyle)
	printPerson(Stan)
	printPerson(Kenny)

}

type PersonI interface {
	GetName() string
	GetAge() int
}

type Student struct {
	Name string
	Age  int
}

func (stu Student) GetName() string {
	return stu.Name
}

func (stu Student) GetAge() int {
	return stu.Age
}

func printPerson(p PersonI) {
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
}
