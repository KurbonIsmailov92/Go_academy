package repository

import (
	"time"
	"todo-crud/internals/db"
	"todo-crud/internals/models"
	"todo-crud/logger"
)

func GetAllUsers() (users []models.User, err error) {
	users = []models.User{}
	err = db.GetDBConnection().Order("id").Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository] Error getting all users: %v", err)
		return nil, err

	}

	return users, nil
}

func GetUserByID(id int) (user models.User, err error) {
	user = models.User{}
	err = db.GetDBConnection().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository] Error getting user: %v", err)
		return user, err
	}

	return user, nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	user = models.User{}
	err = db.GetDBConnection().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository] Error getting user: %v", err)
		return user, err
	}

	return user, nil
}

func CreateNewUser(user models.User) error {
	if err := db.GetDBConnection().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository] Error creating new user: %v", err)
		return err
	}

	return nil
}

func UpdateUser(id int, user, existUser models.User) error {
	if err := db.GetDBConnection().Model(&existUser).
		Updates(user).
		Where("id = ?", id).Error; err != nil {
		logger.Error.Printf("[repository] Error updating user: %v", err)
		return err
	}

	return nil
}

func DeActiveUserByID(id int) error {
	if err := db.GetDBConnection().Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_active":      false,
			"deactivated_at": time.Now(),
		}).Error; err != nil {
		logger.Error.Printf("[repository] Error deactivating user: %v", err)
		return err
	}
	return nil
}
func ActiveUserByID(id int) error {
	if err := db.GetDBConnection().Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_active":      true,
			"deactivated_at": nil, //"0001-01-01 00:00:00",
		}).Error; err != nil {
		logger.Error.Printf("[repository] Error activating user: %v", err)
		return err
	}
	return nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConnection().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository] error getting user by username and password: %v\n", err)
		return user, err
	}

	return user, nil
}
