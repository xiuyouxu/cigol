package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func RestTest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	name := vars["name"]
	id := r.FormValue("id")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "hello,", name, ", your id is", id)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", RestTest).Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("D:/workspace/truecloudstatic"))))

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
