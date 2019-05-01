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
	fmt.Print("Finding All Users: ")
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// Find One users
func (m *UserDatabase) FindOne(id bson.ObjectId) (User, error) {
	fmt.Print("Finding One User: ")
	fmt.Print(id)
	var users User
	err := db.C(COLLECTION).Find(bson.M{"_id":id}).One(&users)
	return users, err
}

// Create a User
func (m *UserDatabase) CreateUser(user User) (User, error) {
	fmt.Print("Creating User: ")
	//fmt.Print(user.id)
	err := db.C(COLLECTION).Insert(user)
	return user, err
}

// Update a User
func (m *UserDatabase) UpdateUser(user User) (User, error) {
	fmt.Print("Update User: ")
	//fmt.Print(user.id)
	err := db.C(COLLECTION).Update(bson.M{"_id":user.ID},user);
	return user, err
}

// Delete a User
func (m *UserDatabase) DeleteUser(id bson.ObjectId) (string, error) {
	fmt.Print("Deleting One User: ")
	fmt.Print(id.String())
	err := db.C(COLLECTION).Remove(bson.M{"_id":id});
	return "success",err;
}