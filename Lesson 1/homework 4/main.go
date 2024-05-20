package main

import (
	"fmt"
	"strings"
)

// 4. Приветствие пользователя:
// Задача: Напишите программу, которая запрашивает у пользователя его имя, а
// затем выводит приветственное сообщение с использованием этого имени.

func main() {
	var userName string

	fmt.Println("Введите свое имя:")
	fmt.Scan(&userName)

	userName = strings.Title(strings.ToLower(strings.TrimSpace(userName)))

	fmt.Printf("Салом! Добро пожаловать, %s!", userName)

}
