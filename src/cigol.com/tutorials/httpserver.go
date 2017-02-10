package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// by default, the params are not to be parsed
	r.ParseForm()
	w.Header().Set("myheader", "good")
	io.WriteString(w, "Hello, world!")
	fmt.Printf(r.URL.Path + "\n")
	//	fmt.Println("param a:", r.Form["a"])
	for k, v := range r.Form {
		fmt.Println("key:", k, "val:", v)
	}
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
