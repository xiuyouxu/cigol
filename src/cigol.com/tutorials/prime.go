package main

import (
	"fmt"
	"math"
)

func IsPrime(n int64) bool {
	m:=int64(math.Sqrt(float64(n)))
	var i int64=2
	for ;i<=m;i++ {
		if n%i==0 {
			return false
		}
	}
	return true
}

/**
	令p1=2,p2=3,p3=5,……表示所有素数，计算p1*p2*...*pn+1是否为素数
*/
func main() {
	var n int64=10000
	var p [5000]int64
	var h int=0
	var i int64=2
	for ;i<n;i++ {
		if IsPrime(i) {
			p[h]=i
			h++
		}
	}
	fmt.Println(p)

	for i:=0;i<h;i++ {
		var prod int64=1
		for j:=0;j<=i;j++ {
			prod*=p[j]
		}
		if IsPrime(prod+1) {
			fmt.Println(i, " prod is prime: ", prod+1)
		} else {
			fmt.Println(i, " prod is not prime: ", prod+1)
		}
	}
}
