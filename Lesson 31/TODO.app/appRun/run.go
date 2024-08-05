package appRun

import (
	"fmt"
	"todo_app/controllers"
)

var commandID int
var TaskMenu = `1 - Создать задачу
2 - Показать список задач
3 - Изменить задачу
4 - Отметить задачу как сделанную
5 - Отметить задачу как НЕсделанную
6 - Удалить задачу из списка
7 - Вернуть задачу в список
0 - Закрыть приложение`

func Run() {
	fmt.Println("Добро пожаловать в приложение TODO.app")
	for {
		fmt.Println(TaskMenu)
		_, err := fmt.Scanln(&commandID)
		if err != nil {
			fmt.Println(err)
		}
		switch commandID {
		case 1:
			controllers.CreateNewTask()
		case 2:
			controllers.GetAllTasks()
		case 3:
			controllers.GetAllTasks()
			fmt.Println("Введите ID задачи:")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
			}
			controllers.GetTaskByID(id)
			controllers.UpdateTask(id)
		case 4:
			controllers.GetAllTasks()
			fmt.Println("Введите ID задачи:")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
			}
			controllers.DoneTask(id)
		case 5:
			controllers.ShowDoneTasks()
			fmt.Println("Введите ID задачи:")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
			}
			controllers.UndoneTask(id)
		case 6:
			controllers.GetAllTasks()
			fmt.Println("Введите ID задачи:")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
			}
			controllers.DeleteTask(id)
		case 7:
			controllers.ShowDeletedTasks()
			fmt.Println("Введите ID задачи:")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
			}
			controllers.CancelTaskDeletion(id)
		case 0:
			fmt.Println("Всего доброго! Программа закрыта")
			return
		default:
			fmt.Println("Введена неверная команда!")
		}
	}
}
