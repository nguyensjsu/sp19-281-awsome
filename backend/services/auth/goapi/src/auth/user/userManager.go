package user

import (
	"../db"
	"../models"
	"../utils"
)

func ValidateUser(user models.User) (bool, models.User) {

	dbUser, userExist := db.GetUserByEmail(user.Email)
	if !userExist {
		return false, models.User{}
	}

	validUser := utils.CompareHash(user.Password, dbUser.Password)

	// mask password
	dbUser.Password = "***"

	return validUser, dbUser
}

func AddUser(user models.User) bool {
	return db.AddUser(models.User{
		Email: user.Email,
		Password: utils.GetPasswordHash(user.Password),
		Role: user.Role,
		Name:user.Name})
}

func AddSession(user models.User) string {
	return db.AddToken(user)
}

func InvalidateSession(token string) {
	db.DeleteToken(token)
}

//func GetSession(token string) (models.Session,bool) {
//	return db.GetToken(token)
//}
