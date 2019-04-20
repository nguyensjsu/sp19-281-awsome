package DAO

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	. "users/models"
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

// Find list of users
func (m *UserDatabase) FindAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// Create a User
func (m *UserDatabase) CreateUser(user User) (User, error) {
	fmt.Print("Creating User: ")
	fmt.Print(user)
	err := db.C(COLLECTION).Insert(user)
	return user, err
}
