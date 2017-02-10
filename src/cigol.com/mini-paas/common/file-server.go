package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starting file server...")
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("html"))))

	srv := &http.Server{
		Handler:      r,
		Addr:         ":81",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
