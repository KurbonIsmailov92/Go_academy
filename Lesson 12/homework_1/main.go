package main

import (
	"bufio"
	"fmt"
	"os"
)

// Задача 1:
// Определить структуру Book с полями Title, Author и Pages. Создайте два
// экземпляра этой структуры, инициализируйте их значениями и выведите информацию о
// книгах на консоль.

func main() {
	reader := bufio.NewReader(os.Stdin)
	var firstBook, secondBook Book

	fmt.Println("Введите информацию о первой книге!")
	firstBook = *firstBook.NewBookMaker(reader)
	fmt.Print("Первая книга внесена!\n")

	fmt.Println("Введите информацию о первой книге!")
	secondBook = *secondBook.NewBookMaker(reader)
	fmt.Println("Вторая книга внесена!")

	fmt.Println("Итого:")
	firstBook.BookInfoPrinter()
	secondBook.BookInfoPrinter()

}

// Структура Книги
type Book struct {
	Title  string
	Author string
	Pages  uint
}

// Создание экземпляра структуры
func (book *Book) NewBookMaker(reader *bufio.Reader) *Book {

	title := book.NewBookTitleSeter(reader)
	author := book.NewBookAuthorSeter(reader)
	pages := book.NewBookCountOfPagesSeter(reader)

	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

// Внесение названия книги
func (book *Book) NewBookTitleSeter(reader *bufio.Reader) (title string) {

	fmt.Println("Введите название книги:")
	title, _ = reader.ReadString('\n')

	return
}

// Внесение автора книги
func (book *Book) NewBookAuthorSeter(reader *bufio.Reader) (author string) {

	fmt.Println("Введите автора книги:")
	author, _ = reader.ReadString('\n')

	return
}

// Внесение количества страниц
func (book *Book) NewBookCountOfPagesSeter(reader *bufio.Reader) (pages uint) {

	fmt.Println("Введите количество страниц книги:")
	fmt.Scan(&pages)
	reader.ReadString('\n')

	return

}

// Вывод информации о книгах
func (book Book) BookInfoPrinter() {
	fmt.Println("_____________________________________________________________________________________")
	fmt.Printf("\nНазвание: %s\n Автор: %s\n Кол-во страниц: %d\n", book.Title, book.Author, book.Pages)
	fmt.Println("_____________________________________________________________________________________")
}
