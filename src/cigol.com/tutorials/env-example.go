package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	os.Setenv("FOO","1")
	
	for index,e:=range os.Environ() {
		fmt.Println(index)

		pair:=strings.Split(e,"=")
		fmt.Println(pair[0])
	}
}