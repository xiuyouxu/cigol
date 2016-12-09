package main

import (
	"log"
	"net/http"
)

func main() {
	d := http.Dir("d:/workspace/go-tutorial/go-crawler/pages/static")
	handler := http.FileServer(d)
	http.Handle("/static/", http.StripPrefix("/static", handler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
