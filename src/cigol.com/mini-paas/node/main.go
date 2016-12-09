package node

import (
	"common"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var config common.Config = common.ReadConfig("./conf.ini")

func Start() {
	fmt.Println("starting node...")
	wait := make(chan int)
	go register(wait)
	startServer(wait)

	fmt.Println("node started...")
}

func register(wait chan int) {
	timeout := make(chan int)
	defer func() {
		time.Sleep(time.Second * 5)
		timeout <- 0
	}()
	select {
	case <-timeout:
		fmt.Println("register...")
	case <-wait:
		fmt.Println("starting failed, cancel registering")
	}
}

func startServer(wait chan int) {
	http.HandleFunc("/op", opHandler)

	port := config["server.port"]
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		wait <- 0
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func opHandler(w http.ResponseWriter, r *http.Request) {
	// by default, the params are not to be parsed
	r.ParseForm()
	//	r.Body.Read()
	op := r.FormValue("op")
	message := r.FormValue("message")

	ret := exec(op, message)
	io.WriteString(w, ret)
	//	fmt.Printf(r.URL.Path + "\n")
	//	fmt.Println("param a:", r.Form["a"])
	//	for k, v := range r.Form {
	//		fmt.Println("key:", k, "val:", v)
	//	}
}

func exec(op, message string) string {
	fmt.Println(op, message)
	var s string
	switch op {
	case "DEPLOY":
		s = doDeploy(message)
	case "START":
		s = doStart(message)
	case "STOP":
		s = doStop(message)
	case "DESTROY":
		s = doDestroy(message)
	default:
		s = common.WrapMessage("result", false, "message", "unrecognized op found: "+op)
	}
	return s
}

func doDeploy(message string) string {
	var msg map[string]interface{}
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		return common.WrapMessage("result", false, "message", err)
	}
	image := msg["image"]
	replicas := msg["replicas"]
	instanceCode := msg["instanceCode"]
	fmt.Println(image, replicas, instanceCode)
	return common.WrapMessage("result", true, "message", "deploy successfully")
}

func doStart(message string) string {
	return common.WrapMessage("result", true, "message", "start successfully")
}

func doStop(message string) string {
	return common.WrapMessage("result", true, "message", "stop successfully")
}

func doDestroy(message string) string {
	return common.WrapMessage("result", true, "message", "destroy successfully")
}
