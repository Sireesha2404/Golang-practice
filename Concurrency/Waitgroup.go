package main

import (
	"fmt"
)

func main() {
	foo()
	bar()
}
func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
