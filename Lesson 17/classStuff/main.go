package main

import "fmt"

func main() {
	d := Dog{
		name:  "Rex",
		sound: "Woof",
	}

	c := Cat{
		name:  "Murka",
		sound: "Meow",
	}

	AnimalPerformer(d)
	AnimalPerformer(c)

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
