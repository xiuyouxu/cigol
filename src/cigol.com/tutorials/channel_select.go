package main

import (
	"fmt"
	"time"
	"math"
)

func fibonacci(c,quit chan int) {
	x,y:=1,1
	for {
		select {
			case c<-x:
			x,y=y,x+y
			fmt.Println("\t",float32(y)/float32(x))
			case <-quit:
			fmt.Println("quit")
			return
			case <-time.After(5*time.Second):
			fmt.Println("timeout")
			break
//			default:
//			fmt.Println("no event...")
		}
	}
}

func main() {
	fmt.Println((math.Sqrt(5)+1)/float64(2))
	c:=make(chan int)
	quit:=make(chan int)
	go func(){
		for {
			fmt.Println(<-c)
			time.Sleep(3*time.Second)
		}
		quit<-0
	}()
	fibonacci(c,quit)
}
