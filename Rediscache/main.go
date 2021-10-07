package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("name", "tilliot", 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

	json, err := json.Marshal(Author{Name: "Tilliot", Age: 26})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id4568", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	value, err := client.Get("id4568").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)
}
