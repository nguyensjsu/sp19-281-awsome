package models

type User struct {
	//ID  bson.ObjectId `bson:"_id" json:"id"`
	Email string        `bson:"email" json:"email"`
	FirstName string	`bson:"firstName" json:"firstName"`
	LastName string	`bson:"lastName" json:"lastName"`
	Profile string	`bson:"profile" json:"profile"`
}

type Login struct {
	Email string        `bson:"email" json:"email"`
	Password string	`bson:"password" json:"password"`
}
