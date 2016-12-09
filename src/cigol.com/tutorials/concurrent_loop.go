package main

import (
	"fmt"
	"strconv"
)

func toString(a interface{}) string{
     if  v,p:=a.(int);p{
        return strconv.Itoa(v)
     }

    if v,p:=a.(float64);  p{
     return strconv.FormatFloat(v,'f', -1, 64)
    }

    if v,p:=a.(float32); p {
        return strconv.FormatFloat(float64(v),'f', -1, 32)
    }

     if v,p:=a.(int16); p {
        return strconv.Itoa(int(v))
     }
      if v,p:=a.(uint); p {
        return strconv.Itoa(int(v))
     }
      if v,p:=a.(int32); p {
        return strconv.Itoa(int(v))
     }
    return "wrong"
}

func main() {
	N:=20
	sem:=make(chan int,N)
	for i:=0;i<N;i++ {
		go func(i int) {
			// do something
			fmt.Println("loop " + toString(i))
			sem<-0;
		}(i)
	}

	for i:=0;i<N;i++ {
		<-sem
	}
}
