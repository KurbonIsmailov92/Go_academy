package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/tasks/", GetAllTasks)
	router.POST("/tasks/", NewTask)
	router.PUT("/tasks/:id", UpdateTask)
	router.DELETE("/tasks/:id", DeleteTask)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}

type Task struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	TaskPriorityID int       `json:"task_priority_id"`
	UserName       string    `json:"user_name"`
	CreatedAt      time.Time `json:"created_at"`
	IsDone         bool      `json:"is_done"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsDeleted      bool      `json:"is_deleted"`
	DeletedAt      time.Time `json:"deleted_at"`
}

var Tasks = []Task{
	{ID: 1,
		Title:          "Приветствие",
		Description:    "Сказать Привет",
		TaskPriorityID: 2,
		UserName:       "Vasya",
		CreatedAt:      time.Now(),
		IsDone:         false,
		UpdatedAt:      time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
		IsDeleted:      false,
		DeletedAt:      time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:             2,
		Title:          "Представиться",
		Description:    "Назвать свое имя",
		TaskPriorityID: 2,
		UserName:       "Petya",
		CreatedAt:      time.Now(),
		IsDone:         false,
		UpdatedAt:      time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
		IsDeleted:      false,
		DeletedAt:      time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
	},
}

func GetAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": Tasks})
}

func NewTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	task.CreatedAt = time.Now()
	Tasks = append(Tasks, task)
	c.JSON(http.StatusOK, "Task created")
}

func UpdateTask(c *gin.Context) {
	var updTask Task

	id, err := strconv.Atoi(c.Param("id"))
	if err = c.BindJSON(&updTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i, task := range Tasks {
		if task.ID == id {
			updTask.ID = id
			updTask.UpdatedAt = time.Now()
			Tasks[i] = updTask
			c.JSON(http.StatusOK, "Task updated")
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTask(c *gin.Context) {
	var deleteTask Task
	id, err := strconv.Atoi(c.Param("id"))
	if err = c.BindJSON(&deleteTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i, task := range Tasks {
		if task.ID == id {
			task.ID = id
			Tasks[i].IsDeleted = deleteTask.IsDeleted
			Tasks[i].DeletedAt = time.Now()
			Tasks[i].UpdatedAt = time.Now()

			c.JSON(http.StatusOK, "Task deleted")
			return
		}

	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
