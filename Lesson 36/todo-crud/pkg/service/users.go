package service

import (
	"errors"
	"gorm.io/gorm"
	"todo-crud/internals/models"
	"todo-crud/pkg/repository"
	"todo-crud/utils"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id int) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateNewUser(user models.User) error {
	_, err := repository.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	user.Password = utils.GenerateHash(user.Password)

	if err = repository.CreateNewUser(user); err != nil {
		return err
	}
	return nil
}

func UpdateUser(id int, user models.User) error {
	existUser, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}
	user.Username = existUser.Username
	user.Password = existUser.Password

	if err = repository.UpdateUser(id, user, existUser); err != nil {
		return err
	}
	return nil
}

func DeActiveUser(id int) error {
	if err := repository.DeActiveUserByID(id); err != nil {
		return err
	}
	return nil
}

func ActivateUser(id int) error {
	if err := repository.ActiveUserByID(id); err != nil {
		return err
	}
	return nil
}
