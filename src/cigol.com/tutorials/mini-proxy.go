package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if args == nil || len(args) < 2 {
		fmt.Println("Usage: mini-proxy <ip> <port>")
		os.Exit(1)
	}
	addr := args[0] + ":" + args[1]
	fmt.Println("Proxying", addr)

	port := 65535
	fmt.Println("Listening port:", port)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error listening on port", port)
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		in, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting on port", port)
			fmt.Println(err)
			continue
		}
		fmt.Println("Accept connection from:", in.RemoteAddr())
		go doProxy(in, addr)
	}
}

func doProxy(in net.Conn, addr string) {
	out, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error connecting", addr)
		fmt.Println(err)
	} else {
		waiter := make(chan int, 2)
		go doTransfer(in, out, waiter)
		go doTransfer(out, in, waiter)
		<-waiter
		<-waiter
		fmt.Println("Finish proxying for ", in.RemoteAddr(), "to", out.RemoteAddr())
	}
}

func doTransfer(in, out net.Conn, waiter chan int) {
	defer in.Close()
	defer out.Close()
	defer func() {
		waiter <- 0
	}()
	var buf [1024]byte
	for {
		n, err := in.Read(buf[0:])
		if err != nil {
			break
		}
		_, err = out.Write(buf[0:n])
		if err != nil {
			break
		}
	}
}
