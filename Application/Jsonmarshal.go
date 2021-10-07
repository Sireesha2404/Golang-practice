package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string
	Age         int
	Active      bool
	lastLoginAt string
}

func main() {
	u, err := json.Marshal(User{Name: "Bob", Age: 10, Active: true, lastLoginAt: "today"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u)) // {"Name":"Bob","Age":10,"Active":true}
}
