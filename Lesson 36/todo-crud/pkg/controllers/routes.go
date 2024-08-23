package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-crud/logger"
)

func RunRoutes() error {
	router := gin.Default()

	router.GET("/ping", PingPong)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	usersG := router.Group("/users")
	{
		usersG.GET("", GetAllUsers)
		usersG.GET(":id", GetUserByID)
		usersG.POST("", CreateNewUser)
		usersG.PUT(":id", UpdateUser)
		usersG.PATCH(":id", DeActivateUser)
		usersG.PATCH(":id/active", ActivateUser)

	}

	tasksG := router.Group("/tasks")
	{
		tasksG.GET("", GetAllTasks)
		tasksG.GET(":id", GetTaskByID)
		tasksG.POST("", CreateNewTask)
		tasksG.PUT(":id", UpdateTaskByID)
		tasksG.GET("done", GetDoneTasks)
		tasksG.GET("deleted", GetDeletedTask)
		tasksG.PATCH(":id", DoneTaskByID)
		tasksG.PATCH(":id/undone", UnDoneTaskByID)
		tasksG.PATCH(":id/delete", DeleteTaskByID)
		tasksG.PATCH(":id/return", CancelTaskDeletion)
	}

	err := router.Run(":8080")
	if err != nil {
		return err
	}

	return nil
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pong": []string{
			"██╗  ██╗███████╗██╗     ██╗      ██████╗ ",
			"██║  ██║██╔════╝██║     ██║     ██╔═══██╗",
			"███████║█████╗  ██║     ██║     ██║   ██║",
			"██╔══██║██╔══╝  ██║     ██║     ██║   ██║",
			"██║  ██║███████╗███████╗███████╗╚██████╔╝",
			"╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝ ╚═════╝ "},
	})
	logger.Info.Println("PingPong said HELLO")
}