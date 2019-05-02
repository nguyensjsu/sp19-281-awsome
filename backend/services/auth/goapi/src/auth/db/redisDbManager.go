package db

import (
	"../models"
	"../utils"
	"fmt"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"time"
)

var redisClient *redis.Client = nil

func ConfigRedis()  {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)

	if err == nil {
		fmt.Println("Connected to Redis")
	}
}

func AddToken(user models.User) string {
	newToken, _ := uuid.NewV4()
	session := models.Session{
		CreatedAt: time.Now(),
		SessionId: newToken.String(),
		Email: user.Email,
		Role: user.Role }

	err := redisClient.Set(newToken.String(), utils.ToGOB64(session), time.Hour).Err()
	if err != nil {
		panic(err)
	}

	return newToken.String()
}

func GetToken(token string) (models.Session, bool) {
	var session models.Session
	sessionExist := true

	cacheValue, err := redisClient.Get(token).Result()
	if err == redis.Nil {
		sessionExist = false
	} else if err != nil {
		panic(err)
	} else {
		session = utils.FromGOB64(cacheValue)
	}
	return session, sessionExist
}

func DeleteToken(token string)  {
	redisClient.Del(token)
}
