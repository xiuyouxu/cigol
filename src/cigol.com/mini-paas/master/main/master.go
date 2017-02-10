package main

import (
	"fmt"

	"cigol.com/mini-paas/master/leadership"
)

func main() {
	fmt.Println("starting master...")
	fmt.Println(leadership.GetLeader())
}
