package main

import (
	"fmt"
)

// apply a big array to cause out of memory error
func main() {
	i := 0
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("========================================")
			fmt.Println(x, "i=", i)
		}
	}()

	a := make([]int, 0)
	for {
		a = append(a, 0)
		i++
	}
	fmt.Println("apply big array successfully", len(a))
}
