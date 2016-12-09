package main

import (
	//"fmt"
)

// send all the integers
func generate(ch chan int) {
	for i:=2;;i++{
		ch<-i
	}
}

// filter with prime p
func filter(in chan int, out chan int, p int) {
	for {
		i:=<-in
		if i%p !=0 {
			out<-i
		}
	}
}

func main() {
	ch:=make(chan int)
	go generate(ch)

	for i:=0;i<30;i++ {
		prime:=<-ch
		print(prime,"\n")
		ch1:=make(chan int)
		go filter(ch, ch1, prime)
		ch=ch1
	}
}
