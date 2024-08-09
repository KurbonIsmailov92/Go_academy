package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
Реализовать 2-endpoint'а:
- Получение списка всех задач
- Создание задачи
*/

func main() {

	http.HandleFunc("/get-all-tasks", GetAllTasks)
	http.HandleFunc("/create-new-task", AddOperation)

	fmt.Println("Сервер поднят! Порт: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

type Task struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	TaskPriorityID int       `json:"task_priority_id"`
	UserName       string    `json:"user_name"`
	IsDone         bool      `json:"is_done"`
	DoneAt         time.Time `json:"done_at"`
	IsDeleted      bool      `json:"is_deleted"`
	DeletedAt      time.Time `json:"deleted_at"`
}

type DefaultResponse struct {
	Message string `json:"message"`
}

var tasks []Task

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	task1 := Task{
		ID:             1,
		Title:          "Приветствие",
		Description:    "Сказать Привет",
		TaskPriorityID: 2,
		UserName:       "Vasya",
		IsDone:         false,
		DoneAt:         time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
		IsDeleted:      false,
		DeletedAt:      time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	tasks = append(tasks, task1)

	task2 := Task{
		ID:             2,
		Title:          "Представиться",
		Description:    "Назвать свое имя",
		TaskPriorityID: 2,
		UserName:       "Petya",
		IsDone:         false,
		DoneAt:         time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
		IsDeleted:      false,
		DeletedAt:      time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	tasks = append(tasks, task2)

	jsonBody, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBody)
	if err != nil {
		panic(err)
	}

}

func AddOperation(w http.ResponseWriter, r *http.Request) {
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		panic(err)
	}

	tasks = append(tasks, task)
	var response DefaultResponse
	response.Message = "Задача добавлена успешно"

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}

}
