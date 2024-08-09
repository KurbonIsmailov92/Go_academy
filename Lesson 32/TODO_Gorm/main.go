package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-gorm/internals/models"
	"todo-gorm/internals/repository"
)

func main() {
	r := gin.Default()

	r.GET("/tasks", func(c *gin.Context) {
		tasks, err := repository.GetAllTasks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tasks)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := repository.SetNewTaskToDB(task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, task)
	})

	// Дополнительные маршруты для других операций...

	r.Run(":8080")
}
