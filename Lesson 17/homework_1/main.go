package main

import "fmt"

/*
Задача 1
Создайте интерфейс Animal с методами MakeSound() и GetName().
Реализуйте структуры Dog и Cat, которые реализуют интерфейс Animal.
Напишите функцию, которая принимает любой объект, реализующий интерфейс "Animal", и вызывает MakeSound и GetName.
*/

func main() {
	rex := Dog{
		name:  "Rex",
		sound: "Woof",
	}

	murka := Cat{
		name:  "Murka",
		sound: "Meow",
	}

	AnimalPerformer(rex)
	AnimalPerformer(murka)

}

type AnimalI interface {
	MakeSound()
	GetAnimalName() string
}

type Dog struct {
	name  string
	sound string
}

func (d Dog) GetAnimalName() string {
	return d.name
}
func (d Dog) MakeSound() {
	fmt.Println(d.sound)
}

type Cat struct {
	name  string
	sound string
}

func (c Cat) GetAnimalName() string {
	return c.name
}
func (c Cat) MakeSound() {
	fmt.Println(c.sound)
}

func AnimalPerformer(a AnimalI) {
	fmt.Println(a.GetAnimalName())
	a.MakeSound()
}
