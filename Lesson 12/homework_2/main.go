package main

import "fmt"

// Задача 2: Определить структуру Rectangle с полями Width и Height. Написать функцию
// Area(), который вычисляет и возвращает площадь прямоугольника. Создать экземпляр
// структуры, вызвать функцию и вывести результат на консоль.

func main() {
	var NewRectangle Rectangle

	NewRectangle = *NewRectangle.NewRectangleMaker()
	fmt.Printf("Прошдь прямоугольника равна %d см²", NewRectangle.Area())

}

type Rectangle struct {
	Width  uint
	Height uint
}

func (r *Rectangle) WidthGeter() (width uint) {
	fmt.Println("Введите длину прямоугольника:")
	fmt.Scan(&width)

	return
}

func (r *Rectangle) HeightGeter() (height uint) {
	fmt.Println("Введите высоту прямоугольника:")
	fmt.Scan(&height)

	return
}

func (r *Rectangle) NewRectangleMaker() *Rectangle {
	width := r.WidthGeter()
	height := r.HeightGeter()

	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

func (r *Rectangle) Area() uint {
	area := r.Width * r.Height

	return area
}
