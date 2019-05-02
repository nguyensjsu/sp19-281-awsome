package db

import (
	"../models"
	"../utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var db *mongo.Database = nil

func ConfigMongoDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(utils.GetMongoUri()))

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

	//defer client.Disconnect(ctx)

	db = client.Database("authdb")

	// create Time To Live index
	//coll := db.Collection("active_sessions")
	//opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	//keys := bsonx.Doc({"createdAt": 1}, { expireAfterSeconds: 60 * 60})
	//index := mongo.IndexModel{}
	//index.Keys = keys
	//_, err = coll.Indexes().CreateOne(context.Background(), index, opts)
	//if err != nil {
	//	log.Fatal(err)
	//}
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

//func AddToken(user models.User) string {
//	newToken, _ := uuid.NewV4()
//	coll := db.Collection("active_sessions")
//	session := models.Session{
//		CreatedAt: time.Now(),
//		SessionId: newToken.String(),
//		Email: user.Email,
//		Role: user.Role }
//	_, err := coll.InsertOne(context.Background(), session)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return newToken.String()
//}

//func GetToken(token string) (models.Session, bool) {
//	var session models.Session
//	sessionExist := true
//	coll := db.Collection("active_sessions")
//	err := coll.FindOne(context.Background(), bson.M{"sessionId": token}).Decode(&session)
//	if err != nil {
//		sessionExist = false
//	}
//	return session, sessionExist
//}
//
//func DeleteToken(token string)  {
//	coll := db.Collection("active_sessions")
//	_, err := coll.DeleteOne(context.Background(), bson.M{"sessionId": token})
//	if err != nil {
//		log.Fatal(err)
//	}
//}