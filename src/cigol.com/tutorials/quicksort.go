/**
* quick sort
* consider sorting an array of integers, we partition the array into 3 parts,
* less than the pivot(always choose the tail of the array as the pivot),
* greater than the pivot and the not dealt part.
* when sorting, try to move the element from the not dealt part to the former two
* until we eliminate the not dealt part. at last, we swap the pivot(the tail) with
* the head of the second part(that is the part with every element being greater
* than the pivot)
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 10
	a := make([]int, count)
	for i := 0; i < count; i++ {
		a[i] = r.Intn(1000)
	}
	fmt.Println(a)

	start := time.Now().UnixNano()
	QuickSort(a)
	end := time.Now().UnixNano()

	fmt.Println("sorting", count, "integers, using time:", end-start, "ns")
	fmt.Println(a)
}

func QuickSort(a []int) {
	qs_helper(a, 0, len(a)-1)
}

func qs_helper(a []int, start, end int) {
	if start < end {
		q := partition(a, start, end)
		qs_helper(a, start, q-1)
		qs_helper(a, q+1, end)
	}
}

func partition(a []int, start, end int) int {
	i := start - 1
	pivot := a[end]
	for j := start; j < end; j++ {
		if a[j] < pivot {
			i++
			if i!=j {
				swap(a, i, j)
			}
		}
	}
	swap(a, i+1, end)
	return i + 1
}

func swap(a []int, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}
