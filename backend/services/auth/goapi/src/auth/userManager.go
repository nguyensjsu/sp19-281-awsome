package main

import (
	"./db"
	"./models"
)

func validateUser(user models.User) bool {

	dbUser, userExist := db.GetUserByEmail(user.Email)
	if !userExist {
		return false
	}

	return compareHash(user.Password, dbUser.Password)
}

func addUser(user models.User) bool {
	return db.AddUser(models.User{Email: user.Email, Password:getPasswordHash(user.Password)})
}

func addSession(email string) string {
	return db.AddToken(email)
}

func invaliateSession(token string) {
	db.DeleteToken(token)
}

func getSession(token string) (string,bool) {
	return db.GetToken(token)
}
