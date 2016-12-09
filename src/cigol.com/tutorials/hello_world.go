package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	fmt.Println("hello world!", runtime.Version())
	// get abs path of current folder
	p, _ := filepath.Abs(".")
	fmt.Println(p)
	fmt.Println(time.Now().UnixNano())

	fmt.Printf("good %q, come %s", "boy", "here")
}
