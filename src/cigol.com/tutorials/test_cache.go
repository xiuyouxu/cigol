package main

import (
	"fmt"
	"math/rand"
	"mathbeta"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("cpu count: ", runtime.NumCPU())

	var cache mathbeta.Xcache = mathbeta.Xcache{make(map[string]interface{}), sync.RWMutex{}}
	cache.Put("one", 101)
	cache.Put("two", 2)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	N := 100000
	waiter := make(chan int, N)
	start := time.Now()
	var result int
	for i := 0; i < N; i++ {
		go func(i int) {
			f := r.Intn(2)
			time.Sleep(time.Duration(f) * time.Second)
			ok := cache.Cas("one", 101, i)
			if ok {
				result = i
			}
			waiter <- 0
		}(i)
	}

	for i := 0; i < N; i++ {
		<-waiter
	}

	end := time.Now()
	fmt.Println("elapsed time: ", end.Sub(start))

	fmt.Println("after concurrent set, one is", cache.Get("one"))
	fmt.Println("cas set one to 3", cache.Cas("one", result, 3))
	fmt.Println("get one", cache.Get("one"))
	fmt.Println("cache size", cache.Getsize())
}
