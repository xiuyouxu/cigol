/* use go routine to generate rand number */
package main

import (
	"fmt"
	"math/rand"
	"time"
)
/* make it look like a "service" to get rand number, as it returns instantly after being called, the rand numbers are put into channel for reading */
func rand_generator() chan int {
	out:=make(chan int)
	go func() {
		r:=rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			out <- r.Intn(100)
		}
	}()
	return out
}

/* multiplex (MUX) version rand number generator */
func mux_rand_generator() chan int {
	// create two rand generators
	generator_1:=rand_generator()
	generator_2:=rand_generator()

	// create a mux channel
	out:=make(chan int)
	// read from generator_1 and integrate to channel out
	go func(){
		for{
			fmt.Println("read from generator 1")
			out<-<-generator_1
		}
	}()
	// read from generator_2 and integrate to channel out
	go  func(){
		for{
			fmt.Println("read from generator 2")
			out<-<-generator_2
		}
	}()
	return out
}

func main() {
	//rand_service_handler:=rand_generator()
	rand_service_handler:=mux_rand_generator()
	for i:=0;i<10;i++ {
		fmt.Printf("%d\n", <-rand_service_handler)
	}
}
