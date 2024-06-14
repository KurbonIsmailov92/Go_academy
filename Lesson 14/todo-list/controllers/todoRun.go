package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func Run() {
	fmt.Println("Добро пожаловать в TODO.app")
	for {
		PrintMenu()
	
		var comandStr string
		reader := bufio.NewReader(os.Stdin)

		fmt.Scan(&comandStr)
		comand, _ := strconv.Atoi(comandStr)
		reader.ReadString('\n')
		
		
		switch comand {
		case 1:
			CreateNewTaskFromConsole()
		case 2:
			ShowTodoList()
		case 3:
			ShowTodoList()
			ChangeNameOfTask()
		case 4:
			ShowTodoList()
			ChangeMemoOfTask()
		case 5:
			ShowTodoList()
			DoneTask()
		case 6:
			ShowTodoList()
			UndoneTask()
		case 7:
			ShowTodoList()
			DeleteTask()
		case 0:
			fmt.Println("Всего доброго! Программа закрыта")
			os.Exit(0)
		default:
			fmt.Println("Вы ввели несуществующую команду.")
		}
	}
}