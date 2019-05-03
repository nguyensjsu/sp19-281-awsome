package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID  bson.ObjectId `json:"id" bson:"_id"`
	Email string `json:"email" bson:"email"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName string `json:"lastname" bson:"lastname"`
	Profile string `json:"profile" bson:"profile"`
	Skills string `json:"skills" bson:"skills"`
	AboutMe string `json:"aboutme" bson:"aboutme"`
	Experience Experience `json:"experience" bson:"experience"`
	Education Education `json:"education" bson:"education"`
}

type Education struct {
	School string `json:"school" bson:"school"`
	Program string `json:"program" bson:"program"`
	Description string `json:"desc" bson:"desc"`
	From int `json:"from" bson:"from"`
	To int `json:"to" bson:"to"`
}

type Experience struct {
	Company string `json:"company" bson:"company"`
	Role string `json:"role" bson:"role"`
	Description string `json:"desc" bson:"desc"`
	From int `json:"from" bson:"from"`
	To int `json:"to" bson:"to"`
}

type Login struct {
	Email string        `bson:"email" json:"email"`
	Password string	`bson:"password" json:"password"`
}