package utils

import (
	"fmt"
	"os"
)

func GetMongoUri() string {
	fmt.Println("MongoUri:", os.Getenv("mongoUri"))
	return os.Getenv("mongoUri")
}

func GetRedisUri() string {
	fmt.Println("RedisUri:", os.Getenv("redisUri"))
	return os.Getenv("redisUri")
}