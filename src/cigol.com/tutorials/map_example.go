package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"1": 2, "pi": 3}
	fmt.Printf("m[1]=%d\n", m["1"])
	fmt.Printf("m[pi]=%d\n", m["pi"])

	fmt.Printf("m[notexist]=%d\n", m["notexist"])
	_, ok := m["notexist"]
	fmt.Println("key 'notexist' exists? ", ok)

	for key, value := range m {
		fmt.Println(key, value)
	}

	for k := range m {
		fmt.Println(k)
	}
}
