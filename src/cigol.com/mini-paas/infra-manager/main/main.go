package main

import (
	"fmt"

	cweb "cigol.com/mini-paas/common/web"
	imweb "cigol.com/mini-paas/infra-manager/web"
)

func main() {
	fmt.Println("starting infra manager...")
	waiter := make(chan int)
	handlers := imweb.GetRestHandlers()
	go cweb.StartServer(waiter, 8080, handlers)
	fmt.Println("infra manager started...")
	<-waiter
	fmt.Println("infra manager exited...")
}
