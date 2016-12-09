package main

import (
	"fmt"
)

func main() {
	i := byte(1 << 3)
	fmt.Println(i)
	j := byte(4)
	fmt.Println(i & j)
}
