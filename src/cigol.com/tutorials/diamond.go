package main

import (
	"fmt"
	"os"
	"strconv"
)

func rowPrint(nspace, nstar int) {
	for j:=0;j<nspace;j++ {
		fmt.Print(" ")
	}
	for j:=0;j<nstar;j++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	args:=os.Args[1:]
	if args==nil || len(args)!=1 {
		fmt.Println("Usage: diamond <num>")
		os.Exit(1)
	}

	n, _:=strconv.Atoi(args[0])
	for i:=1;i<=n;i+=2 {
		k:=(n-i)/2
		rowPrint(k, i)
	}
	for i:=n-2;i>0;i-=2 {
		k:=(n-i)/2
		rowPrint(k, i)
	}
}
