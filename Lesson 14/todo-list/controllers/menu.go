package controllers

import "fmt"

func PrintMenu() {
	fmt.Println("Выберите команду:")
	fmt.Println("1. Добавление новой задачи в список")
	fmt.Println("2. Просмотр всех существующих задач")
	fmt.Println("3. Редактирование название задачи")
	fmt.Println("4. Редактирование описания задачи")
	fmt.Println("5. Отметка задачи как выполненной")
	fmt.Println("6. Отметка задачи как НЕвыполненной")
	fmt.Println("7. Удалить задачу")
	fmt.Println("0. Закрытие программы")
}