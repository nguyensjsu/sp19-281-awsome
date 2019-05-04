package db

import (
	"auth/config"
	"auth/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var db *mongo.Database = nil

func ConfigMongoDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetMongoUri()))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db = client.Database("authdb")

}

func AddUser(user models.User) bool {
	// check if user already exists
	if _, userExist := GetUserByEmail(user.Email); userExist {
		return false
	}

	coll := db.Collection("users")
	_, err := coll.InsertOne(context.Background(),user)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func GetUserByEmail(email string) (models.User, bool) {
	var user models.User
	userExist := true
	coll := db.Collection("users")
	err := coll.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		userExist = false
	}
	return user, userExist
}
