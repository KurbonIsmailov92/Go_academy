package controllers

import (
	"fmt"
	"todo_app/internals/models"
	"todo_app/internals/repository"
)

func CreateNewTask() {
	var task models.Task

	fmt.Println("Введите название задачи:")
	_, err := fmt.Scanln(&task.Title)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Введите описание задачи:")
	_, err = fmt.Scanln(&task.Description)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Введите имя ответственного за задачу:")
	_, err = fmt.Scanln(&task.UserName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Введите приоритет задачи:")
	fmt.Println(`1 - Срочный
2 - Высокий
3 - Средний
4 - Низкий
5 - Бэклог'`)
	_, err = fmt.Scanln(&task.TaskPriorityID)
	if err != nil {
		fmt.Println(err)
	}

	err = repository.SetNewTaskToDB(task)
	if err != nil {
		fmt.Println(err)
	}

}

func GetAllTasks() {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		fmt.Println(err)
	}
	for _, task := range tasks {
		fmt.Print(task, "\n")
	}
}

func GetTaskByID(id int) {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(task)
}

func UpdateTask(id int) {
	fmt.Println("Выберите поле, которую хотите изменить:")
	fmt.Println("1 - Название")
	fmt.Println("2 - Описание")
	fmt.Println("3 - Ответстенного")
	fmt.Println("4 - Приоритет")
	var comandNum int
	_, err := fmt.Scan(&comandNum)
	if err != nil {
		fmt.Println(err)
	}
	switch comandNum {
	case 1:
		fmt.Println("Введите новое название:")
		var newTitle string
		_, err := fmt.Scan(&newTitle)
		if err != nil {
			fmt.Println(err)
		}
		err = repository.UpdateTaskTitleByID(id, newTitle)
		if err != nil {
			fmt.Println(err)
		}
	case 2:
		fmt.Println("Введите новое описание:")
		var newDesc string
		_, err := fmt.Scan(&newDesc)
		if err != nil {
			fmt.Println(err)
		}
		err = repository.UpdateTaskDescriptionByID(id, newDesc)
		if err != nil {
			fmt.Println(err)
		}
	case 3:
		fmt.Println("Введите нового ответственного:")
		var newUsername string
		_, err := fmt.Scan(&newUsername)
		if err != nil {
			fmt.Println(err)
		}
		err = repository.UpdateTaskUserByID(id, newUsername)
		if err != nil {
			fmt.Println(err)
		}
	case 4:
		fmt.Println("Введите новый приоритет:")
		fmt.Println("1 - Срочный, 2 - Высокий, 3 - Средний, 4 - Низкий, 5 - Беклог")
		var newPrID int
		_, err := fmt.Scan(&newPrID)
		if err != nil {
			fmt.Println(err)
		}
		err = repository.UpdateTaskPriorityByID(id, newPrID)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Нет такого поля!")
	}

}

func DoneTask(id int) {
	err := repository.DoneTaskByID(id)
	if err != nil {
		fmt.Println(err)
	}

}

func UndoneTask(id int) {
	err := repository.UnDoneTaskByID(id)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteTask(id int) {
	err := repository.DeleteTaskByID(id)
	if err != nil {
		fmt.Println(err)
	}
}

func CancelTaskDeletion(id int) {
	err := repository.CancelTaskDeletionByID(id)
	if err != nil {
		fmt.Println(err)
	}

}

func ShowDoneTasks() {
	tasks, err := repository.GetDoneTasks()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tasks)
}

func ShowDeletedTasks() {
	tasks, err := repository.GetDeletedTasks()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tasks)
}
