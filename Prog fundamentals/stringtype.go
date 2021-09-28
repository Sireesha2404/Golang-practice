package main

import (
	"fmt"
)

func main() {

	s := "hello india"

	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Printf("%T\n", bs)
	fmt.Println(bs)

}
