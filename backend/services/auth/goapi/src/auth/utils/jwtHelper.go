package utils

import (
	"auth/db"
	"auth/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

func CreateJWT(user models.User) string {

	// Create JWT token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	secretKey, tokenExist := db.GetToken(user.Email)
	if !tokenExist {
		secretKey = db.AddToken(user)
	}
	// claims
	claims := make(jwt.MapClaims)
	claims["sub"] = user.Email
	claims["role"] = user.Role
	claims["firstname"] = user.Name.First
	claims["lastname"] = user.Name.Last
	// Expire in 1 hour
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		panic(err)
	}
	return tokenString
}

func ValidateJWT(inputJWT string) bool {
	if inputJWT == "" {
		return false
	}
	splitToken := strings.Split(inputJWT, "Bearer ")
	if len(splitToken) <= 1 {
		return false
	}
	valid, secret := getClientSecretKey(splitToken[1])
	if !valid {
		return false
	}
	token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return err == nil && token.Valid
}

func InValidateJWT(inputJWT string) bool {
	if inputJWT == "" {
		return false
	}
	splitToken := strings.Split(inputJWT, "Bearer ")
	if len(splitToken) <= 1 {
		return false
	}
	valid, token := getParsedSubject(splitToken[1])
	if valid {
		_, exist := db.GetToken(token.Email)
		if exist {
			db.DeleteToken(token.Email)
			return true
		}
	}
	return false
}

func getParsedSubject(jwt string) (bool, models.Token) {
	splitToken := strings.Split(jwt, ".")
	var token models.Token
	if len(splitToken) == 3 {
		headers := splitToken[1]+"=="
		decodedHeader, err := base64.StdEncoding.DecodeString(headers)
		fmt.Println(string(decodedHeader))
		//if err != nil {
		//	return false, token
		//}
		err = json.Unmarshal(decodedHeader, &token)
		if err != nil {
			return false, token
		} else {
			return true, token
		}
	}
	return false, token
}

func getClientSecretKey(jwt string) (bool, string) {
	valid, token := getParsedSubject(jwt)
	if valid {
		value, sessionValid := db.GetToken(token.Email)
		if sessionValid {
			return true, value
		}
	}
	return false, ""
}
