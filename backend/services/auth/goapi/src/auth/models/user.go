
package models


type User struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}


type Token struct {
	Sub string        `bson:"sub" json:"sub"`
	Role  string        `bson:"role" json:"role"`
}