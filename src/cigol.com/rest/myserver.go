package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// try restful api
func restTest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	name := vars["name"]
	id := r.FormValue("id")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "hello,", name, ", your id is", id)
}

// try html template
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		var m map[string]string = map[string]string{}
		m["token"] = token
		//		fmt.Printf("m[token]=%s\n", m["token"])

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, m)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", restTest).Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("D:/workspace/truecloudstatic"))))
	r.HandleFunc("/login", login).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
