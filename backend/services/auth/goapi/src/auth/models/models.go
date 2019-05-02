
package models

type Name struct {
	First string `bson:"first" json:"first"`
	Last string `bson:"last" json:"last"`
}

type User struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Role string 	`bson:"role" json:"role"`
	Name Name       `bson:"name" json:"name"`
}

type UserDetails struct {
	Name Name `bson:"name" json:"name"`
}

type UserAccount struct {
	Email    string `bson:"email" json:"email"`
	Role string 	`bson:"role" json:"role"`
}

type SignupResponse struct {
	Account UserAccount `bson:"account" json:"account"`
	Details UserDetails `bson:"details" json:"details"`
}

type LoginResponse struct {
	Message string `bson:"message" json:"message"`
	Token string		`bson:"token" json:"token"`
}

type Token struct {
	Email string	`bson:"sub" json:"sub"`
	Role string			`bson:"role" json:"role"`
}