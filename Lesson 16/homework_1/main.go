package main

import "fmt"

/*Задача 1
Создать метод "Area" для вычисления площади фигуры:
Определите структуру Rectangle с полями Width и Height типа float64.
Реализуйте метод Area(), который вычисляет и возвращает площадь прямоугольника.
Создайте объект Rectangle, вызовите метод Area() и выведите результат.*/

func main() {

	var rec Rectangle

	rec.NewRectangleCreator()
	fmt.Printf("Площадь прямоугольника равна %.2f см²", rec.Area())

}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() (result float64) {
	return r.Width * r.Height
}

func (r *Rectangle) NewRectangleCreator() {
	fmt.Println("Введите ширину прямоугольника:")
	fmt.Scan(&r.Width)

	fmt.Println("Введите высоту прямоугольника:")
	fmt.Scan(&r.Height)

}
