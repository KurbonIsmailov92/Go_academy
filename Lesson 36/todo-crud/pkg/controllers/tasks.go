package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-crud/logger"

	"todo-crud/internals/models"
	"todo-crud/pkg/service"
)

func CreateNewTask(c *gin.Context) {
	task := models.Task{}
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Error.Printf("[controller] Failed to bind request body: %v", err)
		return
	}

	if err := service.CreateNewTask(task); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "task created successfully"})
	logger.Info.Printf("[controller] Create new task: %v", task)
}

func GetAllTasks(c *gin.Context) {
	tasks, err := service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	logger.Info.Printf("[controller] Got all tasks: %v", tasks)
}

func GetTaskByID(c *gin.Context) {
	var task models.Task
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Error.Printf("[controllers] Invalid task ID: %v", err)
		return
	}
	task, err = service.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
	logger.Info.Printf("[controller] Got task by id: %v", task)
}

func UpdateTaskByID(c *gin.Context) {
	var task models.Task
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Error.Printf("[controllers] Invalid task ID: %v", err)
		return
	}
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Error.Printf("[controller] Failed to bind task body:%v", err)
		return
	}
	if err = service.UpdateTaskByID(id, task); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
	logger.Info.Printf("[controller] Update task by id: %v", task)
}

func DeleteTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Error.Printf("[controllers] Invalid task ID: %v", err)
		return
	}
	if err = service.DeleteTaskByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
	logger.Info.Printf("[controller] Delete task by id: %v", id)
}

func CancelTaskDeletion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Error.Printf("[controllers] Invalid task ID: %v", err)
		return
	}
	if err = service.CancelTaskDeletionByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task returned successfully"})
	logger.Info.Printf("[controller] Canceled task deletion: %v", id)
}

func DoneTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Error.Printf("[controller] Invalid task ID: %v", err)
		return
	}

	if err := service.DoneTaskByID(id); err != nil {
		logger.Error.Printf("[controllers] Invalid task ID: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task done successfully"})
	logger.Info.Printf("[controller] Done task by id: %v", id)
}

func UnDoneTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid task ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.UnDoneTaskByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task undone successfully"})
}

func GetDoneTasks(c *gin.Context) {
	tasks, err := service.GetDoneTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"done-tasks": tasks})
	logger.Info.Printf("[controller] Got all done tasks: %v", tasks)
}

func GetDeletedTask(c *gin.Context) {
	tasks, err := service.GetDeletedTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted-tasks": tasks})
	logger.Info.Printf("[controller] Got all done tasks: %v", tasks)
}
