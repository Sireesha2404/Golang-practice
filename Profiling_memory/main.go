// main.go
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
		wg.Done()
	}()
	for i := 0; i < 1000; i++ { // running it a 1000 times
		_ = Fib(30)
	}
	wg.Wait()
}

func Fib(n int) uint64 {

	f := make([]uint64, 10000) // make huge allocation

	f[0] = 0
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}
