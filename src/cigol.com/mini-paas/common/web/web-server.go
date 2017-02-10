package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"cigol.com/mini-paas/common/entity"
	"github.com/gorilla/mux"
)

// start the web server
func StartServer(waiter chan int, port int, handlers []entity.RestHandler) {
	defer func() {
		waiter <- 0
	}()
	fmt.Println("starting manager...")

	r := mux.NewRouter()
	for _, handler := range handlers {
		r.HandleFunc(handler.Url, handler.Handler).Methods(handler.Methods...)
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + strconv.Itoa(port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
