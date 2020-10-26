package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// redis Connection options
	options := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // Use default DB
	}

	// DB Connection instance
	db := redis.NewClient(&options)

	// Create Key
	err := db.Set(ctx, "name", "Parul", 0).Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Key = name inserted into Redis")
	}

	// Get Key
	val, err := db.Get(ctx, "name").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key doesn't exist..")
			return
		}
		panic(err)
	}
	fmt.Printf("Key=name has value = %v\n", val)

	// Update Key
	err = db.Set(ctx, "name", "Rahul", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err = db.Get(ctx, "name").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key doesn't exist..")
			return
		}
		panic(err)
	}
	fmt.Printf("After updating, Key=name has value = %v\n", val)

	// Delete key
	err = db.Del(ctx, "name").Err()
	if err != nil {
		panic(err)
	}
	// Search for key to check if deleted successfully
	val, err = db.Get(ctx, "name").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("After Deleting, Key not found...")
		}
	}
	// Increment Counter
	result, err := db.Incr(ctx, "counter").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Counter value after increment = %v\n", result)

	// Decrement Counter
	result, err = db.Decr(ctx, "counter").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Counter Value after Decrement = %v\n", result)
}
