package main

import (
	"fmt"
)

func main() {

	x := map[string]int{
		"James": 24,
		"Bond":  12,
	}
	fmt.Println(x)

	fmt.Println(x["James"])
	fmt.Println(x["Bond"])
}
