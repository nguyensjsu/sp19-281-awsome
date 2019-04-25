// mocked db impl
// todo connect to db
package db

import (
	"../models"
	uuid "github.com/satori/go.uuid"
)

// user password map
var users = map[string]string{}

// token to user map
// todo this should be implemented with ttl
var activeTokens = map[uuid.UUID]string{}

func GetUserByEmail(email string) (models.User, bool) {
	hashedPassword, exist := users[email]
	return models.User{Email: email, Password: hashedPassword}, exist
}

func AddUser(user models.User) bool{
	//check if user exists
	_, userExist := users[user.Email]
	if userExist {
		return false
	} else {
		users[user.Email] = user.Password
		return true
	}
}

func AddToken(email string) string {
	newToken, _ := uuid.NewV4()
	activeTokens[newToken] = email
	return newToken.String()
}

func GetToken(token string) (string, bool) {
	uuidToken, ok := uuid.FromString(token)
	if ok != nil {
		// invalid format token
		panic("invalid token")
	}
	email, tokenExist := activeTokens[uuidToken]
	return email, tokenExist
}

func DeleteToken(token string)  {
	uuidToken, ok := uuid.FromString(token)
	if ok != nil {
		// invalid format token
		panic("invalid token")
	}
	_, tokenExist := activeTokens[uuidToken]
	if tokenExist {
		delete(activeTokens, uuidToken)
	}
}