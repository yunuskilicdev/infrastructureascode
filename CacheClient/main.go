package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-redis/redis"
	"os"
)

func HandleRequest(ctx context.Context) (string, error) {

	redisUrl := os.Getenv("redis_url")
	redisPort := os.Getenv("redis_port")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisUrl, redisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client.Set("1", "1", 0)

	return client.Get("1").String(), nil
}

func main() {
	lambda.Start(HandleRequest)
}
