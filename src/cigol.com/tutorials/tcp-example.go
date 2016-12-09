package main

import (
	"fmt"
	"net"
)

func main() {
	listener, e := net.Listen("tcp", ":80")
	if e != nil {
		fmt.Println(e)
		return
	}
	defer listener.Close()

	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println(e)
			return
		}
		go func() {
			defer conn.Close()
			var b []byte = make([]byte, 2048)
			var l int = -1
			for l, _ = conn.Read(b); l != -1; {
				fmt.Print(string(b[0:l]))
			}
		}()
	}
}
