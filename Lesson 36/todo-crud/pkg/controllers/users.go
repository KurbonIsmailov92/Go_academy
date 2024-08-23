package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-crud/internals/models"
	"todo-crud/logger"
	"todo-crud/pkg/service"
)

func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"massage": "No users found"})
	}
	logger.Info.Printf("[controllers] Successfully got all users: %v", users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserByID(c *gin.Context) {
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Entered wrong id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if user, err = service.GetUserByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"massage": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully got user: %v", user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func CreateNewUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("[controllers] Entered wrond data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	if err := service.CreateNewUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
		return
	}
	logger.Info.Printf("[controllers] Successfully created new user: %v", user)
	c.JSON(http.StatusCreated, gin.H{"Message": "User created successfully!"})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	user.ID = id

	err = service.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"massage": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully updated user: %v", user)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

func DeActivateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
	if err = service.DeActiveUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully deactivated user: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "User deactivated!"})
}

func ActivateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
	if err = service.ActivateUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully activated user: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "User activated!"})
}
