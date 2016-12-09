package main

import (
	"fmt"
	"mathbeta"
)

func main() {
	seeds := []int{7, 11, 13, 17, 19, 23}
	hashes := make([]func(string) int, len(seeds))
	cap := 1024
	for i := 0; i < len(seeds); i++ {
		hashes[i] = mathbeta.GetHash(cap, seeds[i])
	}

	str := []string{"abc", "defg", "hij"}
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i], "============")
		for j := 0; j < len(hashes); j++ {
			fmt.Println(seeds[j], "->", hashes[j](str[i]))
		}
	}

	bits := make([]byte, cap)
	bf := mathbeta.NewBloomFilter(bits, hashes)
	bf.Add("abc")
	bf.Add("defg")
	bf.Add("hij")
	fmt.Println(bf.Contains("abc"))
}
