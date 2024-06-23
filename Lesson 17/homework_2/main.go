package main

import "fmt"

/*
Задача 2:
Создайте интерфейс PrinterI с методом Print(). Затем создайте
структуру Document, которая реализует метод Print. Напишите
функцию printDocument, которая принимает интерфейс PrinterI и
вызывает метод Print.
*/

func main() {

	docx := Document{
		Title: "It`s MS Word document",
		Body:  "This is an example of a MS Word document.",
	}

	pdf := Document{
		Title: "It`s PDF document",
		Body:  "This is an example of a PDF document.",
	}

	printDocument(docx)
	printDocument(pdf)
}

type PrinterI interface {
	Print()
}

type Document struct {
	Title string
	Body  string
}

func (d Document) Print() {
	fmt.Println(d.Title)
	fmt.Println(d.Body)
}

func printDocument(p PrinterI) {
	p.Print()
}
