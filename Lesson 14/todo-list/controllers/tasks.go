package controllers

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"todo/internals/models"
	"todo/internals/repository"
)

func CreateNewTaskFromConsole() {
	var newTask models.Task
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите название задачи:")
	newTask.Name, _ = reader.ReadString('\n')

	fmt.Println("Введите описание задачи:")
	newTask.Memo, _ = reader.ReadString('\n')

	newTask.CreatedAt = time.Now()
	repository.SaveTask(newTask)

}

func ShowTodoList() {
	tasksToShow := repository.GetAllTasks()

	for id, task := range tasksToShow {
		fmt.Println("ID:", id, "Name:", task.Name, "Memo:", task.Memo, "Date:", task.CreatedAt.Format(time.DateOnly), "Is Done:", task.IsDone)
	}

}

func ChangeNameOfTask() {
	var ID uint

	fmt.Println("Введите ID задачи:")
	fmt.Scan(&ID)

	changingTask := repository.GetTaskByID(ID)

	fmt.Println("Введите новое имя задачи:")
	fmt.Scan(&changingTask.Name)

	repository.SaveTaskChanges(ID, changingTask)

}

func ChangeMemoOfTask() {
	var ID uint

	fmt.Println("Введите ID задачи:")
	fmt.Scan(&ID)

	changingTask := repository.GetTaskByID(ID)

	fmt.Println("Введите новое описание задачи:")
	fmt.Scan(&changingTask.Memo)

	repository.SaveTaskChanges(ID, changingTask)

}

func DoneTask() {
	var ID uint

	fmt.Println("Введите ID задачи:")
	fmt.Scan(&ID)

	changingTask := repository.GetTaskByID(ID)
	changingTask.IsDone = true
	repository.SaveTaskChanges(ID, changingTask)
}

func UndoneTask() {
	var ID uint

	fmt.Println("Введите ID задачи:")
	fmt.Scan(&ID)

	changingTask := repository.GetTaskByID(ID)
	changingTask.IsDone = false
	repository.SaveTaskChanges(ID, changingTask)
}

func DeleteTask() {
	var ID uint

	fmt.Println("Введите ID задачи:")
	fmt.Scan(&ID)

	repository.DeleteTaskByID(ID)

}
