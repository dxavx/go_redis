package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/url.insert", ExampleClient)

	}
	err := router.Run(":8080")
	if err != nil {
		return
	}

}

//func ExampleNewClient() {
//	client := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//
//	pong, err := client.Ping().Result()
//	fmt.Println(pong, err)
//	// Output: PONG <nil>
//}

func ExampleClient(*gin.Context) {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Set(ctx, "test_key", "123", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "test_key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("test_key = ", val)

	val2, err := client.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Println()
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
