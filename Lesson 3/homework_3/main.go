package main

import (
	"fmt"
	"strings"
)

// Задача 3: Проверка символа на гласность или согласность
// Напишите программу, которая запрашивает у пользователя один символ и
// проверяет, является ли этот символ гласной или согласной буквой
// алфавита.
// Пример:
// Введите символ: о
// Это гласная буква

func main() {
	var letter string

	fmt.Println("Введите букву:")
	fmt.Scan(&letter)

	letter = strings.ToLower(letter)
	runes := []rune(letter)

	if len(runes) > 1 {

		fmt.Println("Введено несколько символов, Мы взяли только первую", string(runes[0]))
	}

	letter = string(runes[0])

	switch letter {
	case "@", "#", "$", "%", "*", "(", ")", "+", "-", "/", "'", "|", " ", "`", "~", "<", ">", "=", "!", "&", "?", ".", ",":
		fmt.Println("Введена не буква. Это Символ!", letter)
	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
		fmt.Println("Введена не буква. Это цифра!", letter)
	case "a", "u", "o", "e", "i":
		fmt.Println("Это гласная буква", letter)
	case "y":
		fmt.Println("В некоторых момента эта буква считается гласной, как в словах myth, gym", letter)
	case "а", "у", "о", "е", "и", "я", "ю", "ё", "э", "ы":
		fmt.Println("Это гласная буква", letter)
	default:
		fmt.Println("Это согласная буква", letter)
	}

}
