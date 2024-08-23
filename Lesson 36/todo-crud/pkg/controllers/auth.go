package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-crud/internals/models"
	"todo-crud/logger"
	"todo-crud/pkg/service"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Error.Printf("[controller-auth] Failed to bind JSON to struct: %v", err)
		return
	}

	err := service.CreateNewUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("[controller-auth] Successfully created new user: %v", user)
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Error.Printf("[controller-auth] Failed to bind JSON to struct: %v", err)
		return
	}

	accessToken, err := service.SignIn(user.Username, user.Password)
	if err != nil {
		return
	}
	logger.Info.Printf("[controller-auth] Successfully signed in: %v", accessToken)
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
