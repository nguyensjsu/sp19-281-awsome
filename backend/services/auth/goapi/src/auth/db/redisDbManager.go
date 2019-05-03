package db

import (
	"auth/config"
	"auth/models"
	"fmt"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"time"
)

var redisClient *redis.Client = nil

func ConfigRedis()  {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetRedisUri(),
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
	err := redisClient.Set(user.Email, newToken.String(), time.Hour).Err()
	if err != nil {
		panic(err)
	}
	return newToken.String()
}

func GetToken(userEmail string) (string, bool) {
	sessionExist := true

	cacheValue, err := redisClient.Get(userEmail).Result()
	if err == redis.Nil {
		sessionExist = false
	} else if err != nil {
		panic(err)
	}
	return cacheValue, sessionExist
}

func DeleteToken(token string)  {
	redisClient.Del(token)
}
