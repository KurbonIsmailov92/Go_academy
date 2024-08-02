package repository

import (
	"todo_app/internals/db"
	"todo_app/internals/models"
)

func SetNewUserToDB(user *models.User) error {
	_, err := db.GetDBConnection().Exec(db.CreateNewUserQuery, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func GetUserFromDB(user *models.User) *models.User {
	row := db.GetDBConnection().QueryRow(db.FindUserByNameQuery, user.Name)
	row.Scan(&user.ID, &user.Name)
	return user
}

func UpdateUserByID(user *models.User) error {
	_, err := db.GetDBConnection().Exec(db.UpdateUserByIDQuery, user.Name, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByIDFromDB(user *models.User) error {
	_, err := db.GetDBConnection().Exec(db.DeleteUserByIDQuery, user.ID)
	if err != nil {
		return err
	}
	return nil
}
