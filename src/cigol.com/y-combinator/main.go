package main

import (
	"fmt"
	"strings"
)

// original recursive function
func capitalize(s string)string {
	n:=len(s)
	if n==1 {
		return strings.ToUpper(s)
	}
	return capitalize(s[:n-1])+strings.ToUpper(string(s[n-1]))
}

// define the base fn type
type baseFnType func(string)string

// 
type fnType1 func(baseFnType) baseFnType

type fnType2 func(fnType2)baseFnType

func k(f fnType2) baseFnType {
	return func(s string)string{
		n:=len(s)
		if n==1 {
			return strings.ToUpper(s)
		}
		return f(f)(s[:n-1])+strings.ToUpper(string(s[n-1]))
	}
}

func main() {
	upper:=func(f fnType2)baseFnType{
		return f(f)
	}(func(f fnType2)baseFnType{
		return func(s string)string{
			n:=len(s)
			if n==1 {
				return strings.ToUpper(s)
			}
			return f(f)(s[:n-1])+strings.ToUpper(string(s[n-1]))
		}
	})

	kk:=k(k)

	fmt.Println(capitalize("abcdefg"))
	fmt.Println(upper("abcdefg"))
	fmt.Println(kk("abcdefg"))

	// y combinator
	y:=func(fn fnType1)baseFnType{
		return func(f fnType2)baseFnType {
			return f(f)
		}(func(f fnType2)baseFnType {
			return fn(func(s string)string {
				return f(f)(s)
			})
		})
	}
	newupper:=y(func(g baseFnType)baseFnType {
		return func(s string)string {
			n:=len(s)
			if n==1 {
				return strings.ToUpper(s)
			}
			return g(s[:n-1])+strings.ToUpper(string(s[n-1]))
		}
	})
	comma:=y(func(g baseFnType)baseFnType {
		return func(s string)string {
			n:=len(s)
			if n<=3 {
				return s
			}
			return g(s[:n-3])+","+s[n-3:]
		}
	})

	fmt.Println(newupper("abcdefg"))
	fmt.Println(comma("abcdefg"))
}