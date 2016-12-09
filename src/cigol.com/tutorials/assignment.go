package main

import "fmt"

func main(){
	n0,n1:=3,5
	n0,n1=op(n0,n1),n0 // 7,3
	fmt.Printf("n0=%d, n1=%d", n0, n1)
}

func op(n0,n1 int) int {
	return n0*n1-n0-n1
}