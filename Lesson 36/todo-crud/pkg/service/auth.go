package service

import (
	"todo-crud/logger"
	"todo-crud/pkg/repository"
	"todo-crud/utils"
)

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", err
	}

	accessToken, err = GenerateToken(user.ID, user.Username)
	if err != nil {
		logger.Error.Printf("Failed to generate token for user %s: %v", user.Username, err)
		return "", err
	}

	return accessToken, nil
}
