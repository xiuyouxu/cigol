package main

import (
	"fmt"
	"os"
)

func test_panic() {
	var user = os.Getenv("USER1")
	if user=="" {
		panic("no value for $USER1")
	}
}

func main() {
	defer func() {
		if x:=recover();x!=nil {
			fmt.Println(x)
		}
	}()

	test_panic()
	fmt.Println("main func finished...") // not print if test_panic() paniced
}
