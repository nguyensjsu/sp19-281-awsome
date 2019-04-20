package models

type User struct {
	//ID  bson.ObjectId `bson:"_id" json:"id"`
	email string        `bson:"email" json:"email"`
	firstName string	`bson:"firstName" json:"firstName"`
	lastName string	`bson:"lastName" json:"lastName"`
}
