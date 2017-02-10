package main

import (
	"fmt"
	"strconv"

	"cigol.com/mini-paas/common/utils"
	cweb "cigol.com/mini-paas/common/web"
	nweb "cigol.com/mini-paas/node/web"
)

func main() {
	fmt.Println("starting node...")
	waiter := make(chan int)
	handlers := nweb.GetRestHandlers()
	var config utils.Config = utils.ReadConfig("./conf.ini")
	port, _ := strconv.Atoi(config["server.port"])
	go cweb.StartServer(waiter, port, handlers)
	fmt.Println("node started...")
	<-waiter
	fmt.Println("node exited...")
}
