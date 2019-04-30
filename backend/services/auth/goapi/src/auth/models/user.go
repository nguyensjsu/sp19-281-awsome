
package models

import (
	"time"
)

type User struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Role string 	`bson:"role" json:"role"`
}


type Token struct {
	Sub string		`bson:"sub" json:"sub"`
	Role  string	`bson:"role" json:"role"`
}

type Session struct {
	SessionId string	`bson:"sessionId" json:"sessionId"`
	Email string		`bson:"email" json:"email"`
	Role string			`bson:"role" json:"role"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}