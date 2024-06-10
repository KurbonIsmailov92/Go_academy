package main

import "fmt"

// Задача 3: Определить структуру Circle с полями Radius и Color. Написать функцию
// Area() и Circumference(), которые вычисляют площадь и длину окружности
// соответственно. Создать экземпляр структуры, вызвать функцию и вывести результаты
// на консоль.

func main() {

	var myCircle Circle

	myCircle = *myCircle.CircleMaker()
	area := myCircle.Area()
	circumference := myCircle.Circumference()

	fmt.Printf("Площадь круга равна: %.2f Длина окружности равна: %.2f", area, circumference)
}

type Circle struct {
	Radius float64
}

func (c *Circle) CircleMaker() *Circle {
	var radius float64
	fmt.Println("Введите радиус окружности:")
	fmt.Scan(&radius)

	return &Circle{
		Radius: radius,
	}
}

func (c Circle) Area() float64 {
	Pi := 3.14
	area := (c.Radius * c.Radius) * Pi

	return area
}

func (c Circle) Circumference() float64 {
	Pi := 3.14
	circumference := 2 * Pi * c.Radius

	return circumference

}
