package DAO

import (
	"gopkg.in/mgo.v2"
	"log"
)


type UserDatabase struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

func (u *UserDatabase) Connect() {
	session, err := mgo.Dial(u.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(u.Database)
}

