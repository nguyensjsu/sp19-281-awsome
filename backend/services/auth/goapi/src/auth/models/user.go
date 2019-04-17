
package models


type User struct {
	email string        `bson:"email" json:"email"`
	password  string        `bson:"password" json:"password"`
}


type Token struct {
	Sub string        `bson:"sub" json:"sub"`
	Role  string        `bson:"role" json:"role"`
}