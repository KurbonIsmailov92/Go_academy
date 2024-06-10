package main

import "fmt"

type myInt int64

type Person struct {
	Name   string
	Age    myInt
	Salary myInt
	Skills []string
	Hobby map[string]int
}

type Card struct {
	PAN           string
	ExpireDay     string
	CVV           myInt
	CardHolder    string
	AccountNumber string
	Balance       myInt
}

func main() {

	var human Person

	human.Name = "Vasya"
	human.Age = 33
	human.Salary = 2500
	human.Skills = []string{"Kung-Fu", "karate"}
	human.Hobby = map[string]int{"fight": 12}
	

	fmt.Printf("%#v\n", human)
	fmt.Printf("Привет %s! Ваша зарплата равна %d, %v, %v \n", human.Name, human.Salary, human.Skills[0], human.Skills[0])
	for k, v := range human.Hobby {
		fmt.Println(k, v)
	}

}
