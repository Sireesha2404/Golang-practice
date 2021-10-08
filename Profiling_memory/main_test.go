// main_test.go
package main

import "testing"

func TestGetVal(t *testing.T) {
	for i := 0; i < 1000; i++ { // running it a 1000 times
		if Fib(30) != 832040 {
			t.Error("Incorrect!")
		}
	}
}
