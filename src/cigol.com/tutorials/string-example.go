package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	return s.Name + ", " + strconv.Itoa(s.Age)
}

func main() {
	s := &Student{Name: "Jack", Age: 20}
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
