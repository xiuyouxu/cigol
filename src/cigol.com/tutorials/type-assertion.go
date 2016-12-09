package main

import (
	"fmt"
)

func main() {
	var a interface{} = getValue()
	b, ok := a.(map[string]string)
	fmt.Println(b["a"], ok)
}

func getValue() interface{} {
	ret := map[string]string{}
	ret["a"] = "abc"
	return ret
}
