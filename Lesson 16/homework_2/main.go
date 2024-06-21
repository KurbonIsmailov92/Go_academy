package main

import "fmt"

/*
Задача 2
Реализовать метод "Swap" для обмена значениями двух полей:
Определите структуру Pair с полями X и Y типа int.
Реализуйте метод Swap(), который меняет значения полей X и Y местами.
Создайте объект Pair, вызовите метод Swap() и проверьте, что значения полей изменились.
*/

func main() {
	var pair Pair
	pair.NewPairCreator()
	fmt.Printf("X равно: %d Y равно: %d\n", pair.X, pair.Y)

	pair.Swap()
	fmt.Printf("Результат: X равно: %d Y равно: %d", pair.X, pair.Y)

}

type Pair struct {
	X, Y int
}

func (p *Pair) Swap() {
	p.X, p.Y = p.Y, p.X
}

func (p *Pair) NewPairCreator() {
	fmt.Println("Введите Х:")
	fmt.Scan(&p.X)
	fmt.Println("Введите Y:")
	fmt.Scan(&p.Y)
}
