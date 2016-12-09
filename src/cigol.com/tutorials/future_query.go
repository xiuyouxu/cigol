/* Future makes param preparation and function calling decoupling */
package main

import (
	"fmt"
)

type query struct {
	sql chan string
	result chan string
}

func exec_query(q query) {
	go func(){
		sql:=<-q.sql
		// query db and get result
		fmt.Println("query db with: " + sql)
		q.result<- "get " + sql
	}()
}

func main(){
	q:=query{make(chan string,1),make(chan string,1)}
	exec_query(q)

	// add query param
	q.sql<-"select * from user"
	// get result when needed
	fmt.Println(<-q.result)
}
