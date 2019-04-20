package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID  bson.ObjectId `bson:"_id" json:"id"`
	Email string        `bson:"email" json:"email"`
	FirstName string	`bson:"firstName" json:"firstName"`
	LastName string	`bson:"lastName" json:"lastName"`
}
